// File: gpn.go
// 5/4/2014
// Edited by Wangdeqin
// Get number of pages from []byte of a html doc

package gpn

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(0)
	}
}

func GetNumPages(byts []byte) int64 {
	var max_pn int64
	max_pn = -1
	reg := regexp.MustCompile("<div class=\"pagination\">")
	loc1 := reg.FindIndex(byts)
	if loc1 == nil {
		return max_pn
	}
	reg = regexp.MustCompile("</div>")
	loc2 := reg.FindIndex(byts[loc1[1]:])
	reg = regexp.MustCompile("[0-9]+")
	pgnation := byts[loc1[1] : loc1[1]+loc2[0]]
	locs := reg.FindAllIndex(pgnation, -1)
	if locs == nil {
		return max_pn
	}

	var tmp_pn int64
	for i := 0; i < len(locs); i += 1 {
		loc := locs[i]
		tmp_pn, _ = strconv.ParseInt(string((pgnation[loc[0]:loc[1]])), 10, 64)
		if max_pn < tmp_pn {
			max_pn = tmp_pn
		}
	}
	return max_pn
}

