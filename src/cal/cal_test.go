package cal

import (
	"testing"
)

func TestGetDates(t *testing.T) {
	dates := GetDates(2)
	if dates[0] != "2014-05-06" {
		t.Errorf("GetDates: excepted \"2014-05-06\", got \"%s\"", dates[0])
	}
	if dates[1] != "2014-05-05" {
		t.Errorf("GetDates: excepted \"2014-05-05\", got \"%s\"", dates[1])
	}
}
