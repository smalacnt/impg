package tor

import (
    "fmt"
    "net/http"
    "io"
    "os"
)

func GetTor(path, id, urlTmpl string) error {
    url := fmt.Sprintf(urlTmpl, id)
    res, err := http.Get(url)
    // check http.Get
    if err != nil {
        return fmt.Errorf("[DT] %s: %s", id, err.Error())
    }

    // check return content is torrent
    if res.Header["Content-Type"] == nil || res.Header["Content-Type"][0] != "application/x-bittorrent" {
        return fmt.Errorf("[DT] %s: torrent no found url: %s", id, url)
    }

    fileName := id + ".torrent"
    file, err := os.Create(path + "/" + fileName)
    if err != nil {
        return fmt.Errorf("[DT] %s: %s", id, err.Error())
    }
    io.Copy(file, res.Body)
    file.Sync()
    file.Close()
    res.Body.Close()
    return nil
}
