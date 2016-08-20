package gpn

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetNumPages(t *testing.T) {
	file, err := os.Open("pagination.xml")
	defer file.Close()
	if err != nil {
		t.Errorf("can't no open testing file")
	}

	byts, _ := ioutil.ReadAll(file)

	const expected = 519
	actual := GetNumPages(byts)
	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
