package main

import (
    "img"
    "tor"
    "srh"
    "os"
    "fmt"
    "conf"
)

// Search and download imgs/tors
func sit(kwd string, dl_path string, it string) {
    var down func(string, string, string)error
    var urlTmpls []string

    switch it {
    case "img":
        down = img.GetImg
        urlTmpls = conf.IMG_URL_TEMPLATES[:]
    case "tor":
        down = tor.GetTor
        urlTmpls = conf.TOR_URL_TEMPLATES[:]
    default:
        return
    }

    urlLen := len(urlTmpls)

    ids, err := srh.SrhKwd(kwd)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
    println(len(ids), " result(s) for ", kwd)

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

                var urlIndex int
                if t, ok := retry_map[id]; ok {
                    urlIndex = t % urlLen
                } else {
                    urlIndex = 0
                }
                err := down(dl_path, id, urlTmpls[urlIndex])
                fmt.Printf("Downloading %s...\n", id)
                if err != nil {
                    if t, ok := retry_map[id]; ok {
                        if t < conf.RETRY_TIME {
                            go func() {
                                id_chan<- id
                                retry_map[id] += 1
                            } ()
                        } else {
                            fmt.Fprintf(os.Stderr, "%s\n", err.Error())
                        }
                    } else {
                        go func() {
                            id_chan<- id
                            retry_map[id] = 1
                        } ()
                    }
                }
            }
            end_chan <- true
        }(id_chan, end_chan)
    }

    for _, id := range ids {
        id_chan <- id
    }

    for i := 0; i < conf.THREAD_POOL_SIZE; i += 1 {
        id_chan <- "$"
    }

    for i := 0; i < conf.THREAD_POOL_SIZE; i += 1 {
        <-end_chan
    }
}
