/*
Package img implements a function GetImg that takes two arguements,
path and id, and down the image specify by the id to the path.
*/
package img

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetImg down image with id to path
func GetImg(path, id, urlTmpl string) error {
	fileName := path + "/" + id + ".jpg"
	if info, err := os.Stat(fileName); err == nil && info.Size() > 0 {
		return nil
	}
	url := fmt.Sprintf(urlTmpl, id)
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
		return fmt.Errorf("[DI] %s: picture no found url: %s", id, url)
	}

    if res.ContentLength == 0 {
        return fmt.Errorf("[DI] %s: got empty file", id)
    }

	file, err := os.Create(fileName)
	// check create file
	if err != nil {
		return fmt.Errorf("[DI] %s: %s", id, err.Error())
	}
	defer file.Close()
	io.Copy(file, res.Body)
	file.Sync()
	return nil
}
