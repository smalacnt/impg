package tor

import (
	"conf"
	"testing"
)

func TestGetTor(t *testing.T) {
	path := "."
	id := "SPDR006"

	err := GetTor(path, id, conf.TOR_URL_TEMPLATES[0])
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
