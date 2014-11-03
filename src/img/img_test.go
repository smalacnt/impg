package img

import (
    "testing"
)

func TestGetImg(t *testing.T) {
    id := "SPDR006"
    path := "."

    err := GetImg(path, id)

    if err != nil {
        t.Errorf("%s", err.Error())
    }
}
