package main

import (
    "cal"
    "img"
    "tor"
    "srh"
    "fmt"
    "os"
    "conf"
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

    id_chan := make(chan string, conf.ID_CHAN_SIZE)
    end_chan := make(chan bool, conf.THREAD_POOL_SIZE)
    retry_map := make(map[string]int)

    for i := 0; i < conf.THREAD_POOL_SIZE; i += 1 {
        go func (id_chan chan string, end_chan chan bool) {
            for {
                id := <-id_chan
                if id == "$" {
                    break
                }
                err := down(dl_path, id)
                println("Downloading ", id, "...")
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
