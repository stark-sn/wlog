// The reporting package provide means to create work time reports.
package reporting

import (
	"errors"
	"os"
	"s-stark.net/code/wlog/types"
	"text/tabwriter"
	"time"
)

// Create work time report for one day.
func ReportDay(week types.Week, now time.Time) error {
	if week.Days == nil {
		return errors.New("You're not in this week?")
	}

	date := types.Date(now)
	day, ok := week.Days[date]

	if !ok {
		return errors.New("You're not in today?")
	}

	w := tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)

	reportDayOfWeek(w, date, day, now)

	w.Flush()

	return nil
}
