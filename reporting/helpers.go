// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"s-stark.net/code/wlog/types"
	"text/tabwriter"
	"time"
)

// Create work time report for one day.
func reportDayOfWeek(w *tabwriter.Writer, date string, day types.Day, now time.Time) time.Duration {
	dur := sumWorkingTimeDay(day, now)
	breakTime := sumBreakTime(day, now)
	activities, sumActs := sumActivitiesDay(day, now)

	untrackedTime := dur - breakTime - sumActs

	fmt.Fprintf(w, "%s\t\t\n", date)
	fmt.Fprintf(w, "\t\t\n")
	reportSpans(w, day, now)
	fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(dur))
	fmt.Fprintf(w, "\tBreak\t- %s\n", fmtDuration(breakTime))
	dur -= breakTime
	fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(dur))
	fmt.Fprintf(w, "\t\t\n")
	fmt.Fprintf(w, "\tActivities\t\n")
	reportActivities(w, activities)
	fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(sumActs))
	fmt.Fprintf(w, "\tUntracked\t+ %s\n", fmtDuration(untrackedTime))
	fmt.Fprintf(w, "\t\t= %s\n", fmtDuration(dur))

	return dur
}

// Sum up work time of day.
func sumWorkingTimeDay(day types.Day, now time.Time) time.Duration {
	dur := sumSpans(day.GetSpans(now))
	return dur.Round(time.Second)
}

// Sum up break tiem of day.
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
