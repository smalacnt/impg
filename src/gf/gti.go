/*
 * File: gti.go
 * 5/3/2014
 * Get file with filename under a specific path
 */

package gti

import (
	"fmt"
	"ls"
	"tor"
    "os"
)

var fn_path, dl_path string

func getFile(fn_chan chan string, end chan bool) {
	var filename string
	var err error
	for {
		filename = <-fn_chan
		if filename == "$" {
			break
		}
		err = tor.GetFile(dl_path, filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s download fialed\n",
				filename)
		} else {
			fmt.Fprintf(os.Stdout, "OK: %s downloaded\n", filename)
		}
	}
	end <- true
}
func rcv() {
	if x := recover(); x != nil {
		switch x.(type) {
		case string:
			fmt.Fprintf(os.Stderr, "Panic: %s\n", x)
		default:
			panic(x)
		}
	}
}

func GetTorFromImgs(imgs_path string, download_path string) (err error) {
	fn_path = imgs_path
	dl_path = download_path

	/* Get filenames in dir fn_path */
	filenames, err := ls.GetFilenames(fn_path)
	if err != nil {
		return err
	}

	FN_CHAN_SIZE := 10
	fn_chan := make(chan string, FN_CHAN_SIZE)
	end := make(chan bool)

	ROUTINE_NUM := 5
	for i := 0; i < ROUTINE_NUM; i += 1 {
		go getFile(fn_chan, end)
	}
	for _, filename := range filenames {
		if len(filename) < 5 {
			continue
		}
		fn_chan <- filename[:len(filename)-4]
	}
	for i := 0; i < ROUTINE_NUM; i += 1 {
		fn_chan <- "$"
	}

	for i := 0; i < ROUTINE_NUM; i += 1 {
		<-end
	}

    return nil
} /* GetTorFromImgs */
