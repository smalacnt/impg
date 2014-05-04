// File: img.go
// 5/4/2014
// Edited by Wangdeqin
// Download imgs according id to specifc path

package img

import (
    "fmt"
    "net/http"
    "io"
    "os"
)

func GetImg(path string, id string) error {
    url := fmt.Sprintf("http://www.141jav.com/movies/%s.jpg", id)
    res, err := http.Get(url)
    defer func() {
        if res != nil && res.Body != nil {
            res.Body.Close()
        }
    }()

    if err != nil {
        return err
    }

    if res.Header["Content-Type"] == nil || res.Header["Content-Type"][0] != "image/jpeg" {
        return fmt.Errorf("Picture not found")
    }

    file, err := os.Create(path + "/" + id + ".jpg")
    if err != nil {
        return err
    }
    defer file.Close()
    io.Copy(file, res.Body)
    file.Sync()
    return nil
}

