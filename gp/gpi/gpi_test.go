package gpi

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetPageIds(t *testing.T) {
	file, err := os.Open("gpi_test.html")
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	ids := GetPageIds(b)
	if len(ids) < 1 {
		t.Errorf("Got no ids")
	}
	if ids[0] != "SPDR006" {
		t.Errorf("Expceted SPDR006, got %s", ids[0])
	}
}
