package main

import (
    "img"
    "tor"
    "srh"
    "os"
    "fmt"
)

// Search and download imgs/tors
func sit(kwd string, dl_path string, it string) {
    var down func(string, string)error
    switch it {
    case "img":
        down = img.GetImg
    case "tor":
        down = tor.GetTor
    default:
        return
    }

    ids, err := srh.SrhKwd(kwd)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
    println(len(ids), " result(s) for ", kwd)

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
                err := down(dl_path, id)
                println("Downloading ", id, "...")
                if err != nil {
                    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
                }
            }
            end_chan <- true
        }(id_chan, end_chan)
    }

    for _, id := range ids {
        id_chan <- id
    }

    for i := 0; i < THREAD_POOL_SIZE; i += 1 {
        id_chan <- "$"
    }

    for i := 0; i < THREAD_POOL_SIZE; i += 1 {
        <-end_chan
    }
}
