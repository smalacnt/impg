package srh

import (
    "testing"
)

func TestSrhKwd(t *testing.T) {
    kwd := "SW"

    ids, err := SrhKwd(kwd)
    if err != nil {
        t.Errorf("%s", err.Error())
        return
    }

    if len(ids) < 1 {
        t.Errorf("%s", "No search result")
        return
    }

    if ids[0][:2] != "SW" {
        t.Errorf("%s", "Processing search result failed")
        return
    }
}
