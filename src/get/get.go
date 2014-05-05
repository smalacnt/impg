// File: get.go
// 5/4/2014
// Edited by Wangdeqin
// A bunch of imp things

package main

import (
    "srh"
    "tor"
    "img"
    "ls"
    "fmt"
    "os"
    "strings"
)

func usage(progName string) {
    fmt.Printf("Usage: %s <si|st|it> <params>\n", progName)
    fmt.Println("    <si|st> : <kwd> <dl_path>")
    fmt.Println("    <it>    : <img_path> [dl_path]")
    os.Exit(0)
}

func main() {
    if len(os.Args) < 3 {
        usage(os.Args[0])
    }

    switch (os.Args[1]) {
    case "si":
        if len(os.Args) < 4 {
            usage(os.Args[0])
        } else {
            si(os.Args[2], os.Args[3])
        }
    case "st":
        if len(os.Args) < 4 {
            usage(os.Args[0])
        } else {
            st(os.Args[2], os.Args[3])
        }
    case "it":
        var dl_path string
        if len(os.Args) == 3 {
            dl_path = os.Args[2]
        } else {
            dl_path = os.Args[3]
        }
        it(os.Args[2], dl_path)

    default:
        usage(os.Args[0])
    }
}

// Search and download imgs
func si(kwd string, dl_path string) {
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
                err := img.GetImg(dl_path, id)
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

// Search and download torrents
func st(kwd string, dl_path string) {
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
                err := tor.GetTor(dl_path, id)
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
