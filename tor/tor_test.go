package tor

import (
	"conf"
	"testing"
)

func TestGetTor(t *testing.T) {
	path := "."
	id := "SET014"

	err := GetTor(path, id, conf.TOR_URL_TEMPLATES[1])
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
