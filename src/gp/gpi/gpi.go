// File: gpi.go
// 5/4/2014
// Edited by Wangdeqin
// Get ids from []byte of a html doc

package gpi

import (
    "regexp"
)

func GetPageIds(byts []byte) []string {
    reg := regexp.MustCompile("/movies/[A-Z0-9]+\\.jpg")
    locs := reg.FindAllIndex(byts, 20)

    ids := make([]string, 0)
    for i := 0; i < len(locs); i++ {
        loc := locs[i]
        ids = append(ids, string(byts[ loc[0] + 8 : loc[1] - 4 ]))
    }

    return ids
}

