/*
Package img implements a function GetImg that takes two arguements,
path and id, and down the image specify by the id to the path.
*/
package img

import (
    "fmt"
    "net/http"
    "io"
    "os"
)

// GetImg down image with id to path
func GetImg(path string, id string) error {
    url := fmt.Sprintf("http://www.141jav.com/movies/%s.jpg", id)
    res, err := http.Get(url)
    defer func() {
        if res != nil && res.Body != nil {
            res.Body.Close()
        }
    }()

    // check http.Get result
    if err != nil {
        return fmt.Errorf("[DI] %s: %s", id, err.Error())
    }

    // check return content is a img
    if res.Header["Content-Type"] == nil || res.Header["Content-Type"][0] != "image/jpeg" {
        return fmt.Errorf("[DI] %s: picture not found", id)
    }

    file, err := os.Create(path + "/" + id + ".jpg")
    // check create file
    if err != nil {
        return fmt.Errorf("[DI] %s: %s", id, err.Error())
    }
    defer file.Close()
    io.Copy(file, res.Body)
    file.Sync()
    return nil
}
