package img

import (
	"conf"
	"testing"
)

func TestGetImg(t *testing.T) {
	id := "SPDR006"
	path := "."

	err := GetImg(path, id, conf.IMG_URL_TEMPLATES[0])

	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
