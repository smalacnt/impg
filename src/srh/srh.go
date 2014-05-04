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

func SrhKwd(kwd string) (ids []string, err error) {
    ids = make([]string, 0)
    url := fmt.Sprintf("http://www.141jav.com/search/%s/", kwd)

    res, err := http.Get(url)
    if err != nil {
        return
    }
    byts, _ := ioutil.ReadAll(res.Body)
    pgn := gpn.GetNumPages(byts)
    if pgn < 1 {
        err = fmt.Errorf("No search result")
        return
    }

    ids = append(ids, gpi.GetPageIds(byts)...)

    var i int64
    for i = 2; i <= pgn; i += 1 {
        url = fmt.Sprintf("http://www.141jav.com/search/%s/%d/", kwd, i)
        res, err = http.Get(url)
        if err != nil {
            return
        }
        byts, _ = ioutil.ReadAll(res.Body)
        newids := gpi.GetPageIds(byts)
        if len(newids) < 1 {
            break
        }
        ids = append(ids, newids...)
    }

    return
}
