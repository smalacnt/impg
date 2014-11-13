package img

import (
	"conf"
	"testing"
)

func TestGetImg(t *testing.T) {
	id := "BOMN049"
	path := "."

	err := GetImg(path, id, conf.IMG_URL_TEMPLATES[0])

	if err == nil {
		t.Errorf("Expected empty file, but not!")
	}

    id = "SET014"

    err = GetImg(path, id, conf.IMG_URL_TEMPLATES[0])
    if err != nil {
        t.Errorf("Got error: %s", err)
    }

}
