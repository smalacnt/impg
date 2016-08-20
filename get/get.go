/*
Package main implements the main procedure of impg.

The syntax of the user input is:

    $./get <si|st|it> <params>
           <si|st> : <kwd> <dl_path>
           <it>    : <img_path> [dl_path]
           <li>    : <cnt>      [dl_path]
           <lt>    : <cnt>      [dl_path]
*/

package main

import (
	"fmt"
	"github.com/wdeqin/impg/cal"
	"os"
	"strconv"
	"time"
)

func usage(progName string) {
	fmt.Printf("Usage: %s <si|st|it|li|lt> <params>\n", progName)
	fmt.Printf("          <si|st> : <kwd>        [dl_path]\n")
	fmt.Printf("          <it>    : <img_path>   [dl_path]\n")
	fmt.Printf("          <li>    : <cnt>        [dl_path]\n")
	fmt.Printf("          <lt>    : <cnt>        [dl_path]\n")
	fmt.Printf("          <cli>   : <ds>  <de>   [dl_path]\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 3 {
		usage(os.Args[0])
	}

	var dl_path string
	switch os.Args[1] {
	case "si":
		if len(os.Args) < 4 {
			dl_path = "." // working dierctory
		} else {
			dl_path = os.Args[3] // specified by user
		}
		sit(os.Args[2], dl_path, "img")
	case "st":
		if len(os.Args) < 4 {
			dl_path = "." // working dierctory
		} else {
			dl_path = os.Args[3] // specified by user
		}
		sit(os.Args[2], dl_path, "tor")
	case "it":
		var dl_path string
		if len(os.Args) == 3 {
			dl_path = os.Args[2]
		} else {
			dl_path = os.Args[3]
		}
		it(os.Args[2], dl_path)
	case "li":
		if len(os.Args) < 4 {
			dl_path = "." // working dierctory
		} else {
			dl_path = os.Args[3] // specified by user
		}
		i, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wrong <cnt>: %s\n", err.Error())
			os.Exit(1)
		}
		dates := cal.GetDates(int(i))
		lit(dates, dl_path, "img")
	case "lt":
		if len(os.Args) < 4 {
			dl_path = "." // working dierctory
		} else {
			dl_path = os.Args[3] // specified by user
		}
		i, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wrong <cnt>: %s\n", err.Error())
			os.Exit(1)
		}
		dates := cal.GetDates(int(i))
		lit(dates, dl_path, "tor")
	case "cli":
		arg_cnt := len(os.Args)
		if arg_cnt == 4 {
			dl_path = "."
		} else if arg_cnt == 5 {
			dl_path = os.Args[4]
		} else {
			usage(os.Args[0])
			os.Exit(1)
		}
		ds_str := os.Args[2]
		de_str := os.Args[3]
		defer func() {
			if e := recover(); e != nil {
				fmt.Fprintf(os.Stderr, "Wrong date format, eg: 2006-1-2\n")
				os.Exit(1)
			}
		}()
		ds, err := time.Parse("2006-1-2", ds_str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wrong date format, eg: 2006-1-2\n")
			os.Exit(1)
		}
		de, err := time.Parse("2006-1-2", de_str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wrong date format, eg: 2006-1-2\n")
			os.Exit(1)
		}
		dates := cal.GetDatesBtw(ds, de)
		lit(dates, dl_path, "img")
	default:
		usage(os.Args[0])
	}
}
