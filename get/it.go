package main

import (
	"fmt"
	"github.com/wdeqin/impg/conf"
	"github.com/wdeqin/impg/ls"
	"github.com/wdeqin/impg/tor"
	"os"
	"strings"
)

// Get imgs names and download torrents
func it(fn_path string, dl_path string) {
	filenames, err := ls.GetFilenames(fn_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		usage(os.Args[0])
	}

	tornames, err := ls.GetFilenames(dl_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		usage(os.Args[0])
	}

	exist := map[string]bool{}
	for _, torname := range tornames {
		if strings.HasSuffix(torname, ".torrent") {
			exist[torname[:len(torname)-8]] = true
		}
	}

	urlLen := len(conf.TOR_URL_TEMPLATES)
	id_chan := make(chan string, conf.ID_CHAN_SIZE)
	end_chan := make(chan bool, conf.THREAD_POOL_SIZE)
	retry_map := make(map[string]int)

	for i := 0; i < conf.THREAD_POOL_SIZE; i += 1 {
		go func(id_chan chan string, end_chan chan bool) {
			for {
				id := <-id_chan
				if id == "$" {
					break
				}
				if exist[id] {
					continue
				}
				fmt.Printf("Downloading %s...\n", id)
				var urlIndex int
				if t, ok := retry_map[id]; ok {
					urlIndex = t % urlLen
				} else {
					urlIndex = 0
				}
				err := tor.GetTor(dl_path, id, conf.TOR_URL_TEMPLATES[urlIndex])
				if err != nil {
					if t, ok := retry_map[id]; ok {
						if t < conf.RETRY_TIME {
							go func() {
								id_chan <- id
								retry_map[id] += 1
							}()
						} else {
							fmt.Fprintf(os.Stderr, "%s\n", err.Error())
						}
					} else {
						go func() {
							id_chan <- id
							retry_map[id] = 1
						}()
					}
				}
			}
			defer func() {
				end_chan <- true
			}()
		}(id_chan, end_chan)
	}

	go func() {
		for _, filename := range filenames {
			if len(filename) < 5 {
				continue
			}
			if !strings.HasSuffix(filename, ".jpg") {
				continue
			}
			id_chan <- filename[:len(filename)-4]
		}

		for i := 0; i < conf.THREAD_POOL_SIZE; i += 1 {
			id_chan <- "$"
		}
	}()

	for i := 0; i < conf.THREAD_POOL_SIZE; i += 1 {
		<-end_chan
	}
}
