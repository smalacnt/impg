// File: srh.go
// 5/4/2014
// Edited by Wangdeqin
// Search given keyword and return ids

package srh

import (
    "gp/gpi"
    "gp/gpn"
    "net/http"
    "fmt"
    "io/ioutil"
)

func SrhKwd(kwd string) ([]string, error) {
    ids := make([]string, 0)
    url := fmt.Sprintf("http://www.141jav.com/search/%s/", kwd)

    res, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("[S] %s: %s", kwd, err.Error())
    }
    byts, _ := ioutil.ReadAll(res.Body)
    pgn := gpn.GetNumPages(byts)
    ids = append(ids, gpi.GetPageIds(byts)...)
    if pgn < 1 {
        if  len(ids) > 0 {
            return ids, nil
        } else {
            return nil, fmt.Errorf("[S] %s: no search result", kwd)
        }
    }


    const THREAD_POOL_SIZE = 10
    const PGN_CHAN_SIZE = 5
    const ID_CHAN_SIZE = 20
    pgn_chan := make(chan int64, PGN_CHAN_SIZE)
    id_chan := make(chan string, ID_CHAN_SIZE)
    var i int64

    for i := 0; i < THREAD_POOL_SIZE; i += 1 {
        go func(pgn_chan chan int64, id_chan chan string) {
            defer func() {
                id_chan<-"$"
            }()
            for {
                pgn := <-pgn_chan
                if pgn == -1 {
                    break
                }
                url = fmt.Sprintf("http://www.141jav.com/search/%s/%d/", kwd, pgn)
                res, err = http.Get(url)
                if err != nil {
                    continue
                }
                byts, _ = ioutil.ReadAll(res.Body)
                ids := gpi.GetPageIds(byts)
                for _, id := range ids {
                    id_chan <- id
                }
            }
        }(pgn_chan, id_chan)
    }

    go func() {
        for i = 2; i <= pgn; i += 1 {
            pgn_chan <- i
        }

        for i := 0; i < THREAD_POOL_SIZE; i += 1 {
            pgn_chan <- -1
        }
    }()

    end_cnt := 0
    id := ""
    for {
        id = <-id_chan
        if id == "$" {
            end_cnt += 1
            if end_cnt == THREAD_POOL_SIZE {
                break
            }
        } else {
            ids = append(ids, id)
        }
    }
    return ids, nil
}

