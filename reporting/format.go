// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"time"
)

const durationPlaceholder = "             "

func fmtDuration(d time.Duration) string {

	d = d.Round(time.Second)

	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
}

func fmtTime(t time.Time) string {
	return t.Format("15:04:05")
}

func tsFormat(d time.Duration) string {
	return fmt.Sprintf("%05.2f", d.Hours())
}
