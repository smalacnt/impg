package cal

import (
    "time"
	"testing"
)

func TestGetDates(t *testing.T) {
	dates := GetDates(2)
    if len(dates) < 2 {
        t.Errorf("Excepted 2 dates, got %d days\n", len(dates))
    }

    ds, _ := time.Parse("2006-1-2", "2014-04-09")
    de, _ := time.Parse("2006-1-2", "2014-4-11")
    datesBtw := GetDatesBtw(ds, de)
    if len(datesBtw) != 3 {
        t.Errorf("ds = %v\n", ds)
        t.Errorf("de = %v\n", de)
        t.Errorf("GetDates(2014-5-29, 2014-6-4), got %v", datesBtw)
    }
}
