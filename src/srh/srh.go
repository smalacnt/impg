// File: srh.go
// 5/4/2014
// Edited by Wangdeqin
// Search given keyword and return ids

package srh

import (
	"conf"
	"fmt"
	"gp/gpi"
	"gp/gpn"
	"io/ioutil"
	"net/http"
)

func srhPgs(d string, typ string) ([]string, error) {
	var url string
	switch typ {
	case "Lst":
		url = fmt.Sprintf(conf.LATEST_URL_TEMPLATE, d)
	case "Srh":
		url = fmt.Sprintf(conf.SEARCH_URL_TEMPLATE, d)
	default:
		return nil, fmt.Errorf("[SrhPgs] Unknown typ %s", typ)
	}
	println(url)
	ids := make([]string, 0)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("[%s] %s: %s", typ, d, err.Error())
	}
	byts, _ := ioutil.ReadAll(res.Body)
	pgn := gpn.GetNumPages(byts)
	ids = append(ids, gpi.GetPageIds(byts)...)
	if pgn < 1 {
		if len(ids) > 0 {
			return ids, nil
		} else {
			return nil, fmt.Errorf("[%s] %s: no %s result", typ, d, typ)
		}
	}

	const THREAD_POOL_SIZE = 10
	const PGN_CHAN_SIZE = 5
	const ID_CHAN_SIZE = 20
	urlTemplate := url + "/%d"
	pgn_chan := make(chan int64, PGN_CHAN_SIZE)
	id_chan := make(chan string, ID_CHAN_SIZE)
	var i int64

	for i := 0; i < THREAD_POOL_SIZE; i += 1 {
		go func(pgn_chan chan int64, id_chan chan string) {
			defer func() {
				id_chan <- "$"
			}()
			for {
				pgn := <-pgn_chan
				if pgn == -1 {
					break
				}
				pageUrl := fmt.Sprintf(urlTemplate, pgn)
				res, err = http.Get(pageUrl)
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
		for i = 2; i <= pgn; i++ {
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

func SrhLst(date string) ([]string, error) {
	return srhPgs(date, "Lst")
}

func SrhKwd(kwd string) ([]string, error) {
	return srhPgs(kwd, "Srh")
}
