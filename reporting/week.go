// The reporting package provide means to create work time reports.
package reporting

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
	"s-stark.net/code/wlog/types"
)

// Create work time report for week.
func ReportWeek(week types.Week, now time.Time) error {
	if week.Days == nil {
		return errors.New("You're not in this week?")
	}

	var dates []string
	for date, _ := range week.Days {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	var dur time.Duration

	w := tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)

	for _, date := range dates {
		day, _ := week.Days[date]
		dayDur := reportDayOfWeek(w, date, day, now)
		fmt.Fprintf(w, "\t\t%s\n", durationPlaceholder)
		dur += dayDur
	}

	fmt.Fprintf(w, "\t\t%s\n", durationPlaceholder)
	fmt.Fprintf(w, "\t\t= %v\n", fmtDuration(dur))

	w.Flush()

	return nil
}

