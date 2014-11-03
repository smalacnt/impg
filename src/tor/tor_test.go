package tor

import (
    "testing"
)

func TestGetTor(t *testing.T) {
    path := "."
    id := "SPDR006"

    err := GetTor(path, id)
    if err != nil {
        t.Errorf("%s", err.Error())
    }
}
