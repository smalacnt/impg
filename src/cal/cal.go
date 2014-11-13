/*
Package cal generate data string in the form of "2014-05-07"
*/
package cal

import (
	"time"
)

func GetDates(cnt int) []string {
	// take cnt = abs(cnt)
	if cnt < 0 {
		cnt = -cnt
	}

	dates := make([]string, cnt)

	t := time.Now()

	for i := 0; i < cnt; i += 1 {
		t = t.AddDate(0, 0, -1) // backward a day
		dates[i] = t.Format(time.RFC3339)[:10]
	}

	return dates
}

func GetDatesBtw(ds, de time.Time) []string {
    if de.Sub(ds).Hours() <= 0 {
        // swap if not asc
        ds, de = ds, de
    }

    var dates []string
    for de.Sub(ds).Hours() >= 0 {
        dates = append(dates, ds.Format(time.RFC3339)[:10])
        ds = ds.AddDate(0, 0, 1)
    }

    return dates
}
