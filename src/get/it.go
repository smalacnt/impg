package main

import (
    "ls"
    "fmt"
    "os"
    "tor"
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
            exist[torname[:len(torname) - 8]] = true
        }
    }

    const ID_CHAN_SIZE = 10
    const THREAD_POOL_SIZE = 5
    id_chan := make(chan string, ID_CHAN_SIZE)
    end_chan := make(chan bool, THREAD_POOL_SIZE)

    for i := 0; i < THREAD_POOL_SIZE; i += 1 {
        go func(id_chan chan string, end_chan chan bool) {
            for {
                id := <-id_chan
                if id == "$" {
                    break
                }
                if exist[id] {
                    continue
                }

                fmt.Fprintf(os.Stdout, "Downloading %s...\n", id)
                err := tor.GetTor(dl_path, id)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
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
            if !strings.HasSuffix(filename, ".jpg"){
                continue
            }
            id_chan <- filename[:len(filename) - 4]
        }

        for i := 0; i < THREAD_POOL_SIZE; i += 1 {
            id_chan <- "$"
        }
    }()

    for i := 0; i < THREAD_POOL_SIZE; i += 1 {
        <-end_chan
    }
}
