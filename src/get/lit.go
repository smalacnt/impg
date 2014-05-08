package main

import (
    "cal"
    "img"
    "tor"
    "srh"
    "fmt"
    "os"
)

// Latest cnt dates of imgs/tors
func lit(cnt int,  dl_path string, it string) {
    var down func(string, string)error
    switch it {
    case "img":
        down = img.GetImg
    case "tor":
        down = tor.GetTor
    default:
        return
    }

    dates := cal.GetDates(cnt)

    ids := make([]string, 0)

    for _, date := range dates {
        tmp, err := srh.SrhLst(date)
        ids = append(ids, tmp...)
        if err != nil {
            fmt.Fprintf(os.Stderr, "[SrhLst]: %s\n", err.Error())
        }
    }
    println(len(ids), " result(s) for lst ", cnt)

    const ID_CHAN_SIZE = 10
    const THREAD_POOL_SIZE = 5
    id_chan := make(chan string, ID_CHAN_SIZE)
    end_chan := make(chan bool, THREAD_POOL_SIZE)

    for i := 0; i < THREAD_POOL_SIZE; i += 1 {
        go func (id_chan chan string, end_chan chan bool) {
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
