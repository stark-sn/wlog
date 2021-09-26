// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"time"
)

var loc, _ = time.LoadLocation("Europe/Berlin")

func fmtDuration(d time.Duration) string {

	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d.Round(time.Second) / time.Second

	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
}

func fmtTime(t time.Time) string {
	return t.In(loc).Format("15:04:05")
}
