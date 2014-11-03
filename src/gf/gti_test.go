package gti

import (
    "testing"
)

func TestGetTorFromImgs(t *testing.T) {
    imgs_path := "pick"
    download_path := "pick_dl"

    err := GetTorFromImgs(imgs_path, download_path)
    if err != nil {
        t.Errorf("%s", err.Error())
    }
}
