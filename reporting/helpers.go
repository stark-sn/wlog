// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"s-stark.net/code/wlog/types"
)

var writer = tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)

const untracked = "[Untracked]"
const nonWorking = "[Non Working]"

// Create work time report for one day.
func reportDayOfWeek(w io.Writer, date string, day types.Day, now time.Time) time.Duration {
	dur := sumWorkingTimeDay(day, now)
	breakTime := sumBreakTime(day, now)
	nons, nonWorkingTime := sumNonWorkingTime(day)
	activities, sumActs := sumActivitiesDay(day, now)

	untrackedTime := dur - breakTime - sumActs
	dur = dur + nonWorkingTime

	fmt.Fprintf(w, "%s\t\t\n", date)
	fmt.Fprintf(w, "\t\t\n")

	if nonWorkingTime > 0 {
		fmt.Fprintf(w, "\t%s\t\n", nonWorking)
		reportNonWorkingTime(w, nons)
		fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(nonWorkingTime))
	}

	reportSpans(w, day, now)
	fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(dur))

	if breakTime > 0 {
		fmt.Fprintf(w, "\tBreak\t- %s\n", fmtDuration(breakTime))
		dur -= breakTime
		fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(dur))
	}

	fmt.Fprintf(w, "\t\t\n")

	if len(activities) > 0 {
		fmt.Fprintf(w, "\t[Activities]\t\n")
		reportActivities(w, activities)
		fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(sumActs))
	}

	if nonWorkingTime > 0 {
		fmt.Fprintf(w, "\t%s\t+ %s\n", nonWorking, fmtDuration(nonWorkingTime))
	}

	if untrackedTime > 0 {
		fmt.Fprintf(w, "\t%s\t+ %s\n", untracked, fmtDuration(untrackedTime))
	}

	fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(dur))

	return dur
}

// Sum up work time of day.
func sumWorkingTimeDay(day types.Day, now time.Time) time.Duration {
	dur := sumSpans(day.GetSpans(now))
	return dur.Round(time.Second)
}

// Sum up break time of day.
func sumBreakTime(day types.Day, now time.Time) time.Duration {
	dur := sumSpans(day.GetBreaks(now))
	return dur.Round(time.Second)
}

// Sum activities of day.
func sumActivitiesDay(day types.Day, now time.Time) (map[string]time.Duration, time.Duration) {
	var dur time.Duration
	var durs = make(map[string]time.Duration)

	for _, act := range day.GetActivities(now) {
		d := act.Duration()

		durs[act.Title] += d
		dur += d
	}

	return durs, dur.Round(time.Second)
}

// Sum time spans.
func sumSpans(spans []types.Span) time.Duration {
	var dur time.Duration

	for _, span := range spans {
		dur += span.Duration()
	}

	return dur.Round(time.Second)
}

// Sum non working time of day.
func sumNonWorkingTime(day types.Day) (map[string]time.Duration, time.Duration) {
	var dur time.Duration
	var durs = make(map[string]time.Duration)

	for _, non := range day.NonWorkingTime {
		dur += non.Duration
		durs[non.Title] += non.Duration
	}

	return durs, dur.Round(time.Second)
}
