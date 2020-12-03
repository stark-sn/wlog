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

	slackTime := dur - breakTime - sumActs

	fmt.Fprintf(w, "%v\t\t%s\n", date, durationPlaceholder)
	fmt.Fprintf(w, "\t\t%s\n", durationPlaceholder)
	reportSpans(w, day)
	fmt.Fprintf(w, "\t\t= %v\n", fmtDuration(dur))
	fmt.Fprintf(w, "\tBreak\t- %v\n", fmtDuration(breakTime))
	dur -= breakTime
	fmt.Fprintf(w, "\t\t= %v\n", fmtDuration(dur))
	fmt.Fprintf(w, "\t\t%s\n", durationPlaceholder)
	fmt.Fprintf(w, "\tActivities\t%s\n", durationPlaceholder)
	reportActivities(w, activities)
	fmt.Fprintf(w, "\t\t= %v\n", fmtDuration(sumActs))
	fmt.Fprintf(w, "\tSlack\t+ %v\n", fmtDuration(slackTime))
	fmt.Fprintf(w, "\t\t= %v\n", fmtDuration(dur))

	return dur
}

// Sum up work time of day.
func sumWorkingTimeDay(day types.Day, now time.Time) time.Duration {
	dur := sumSpans(day.Spans)

	if day.IsIn() {
		dur += now.Sub(day.CurSpan.Start)
	}

	return dur.Round(time.Second)
}

// Sum up break tiem of day.
func sumBreakTime(day types.Day, now time.Time) time.Duration {
	dur := sumSpans(day.Breaks)

	if day.IsOnBreak() {
		dur += now.Sub(day.CurBreak.Start)
	}

	return dur.Round(time.Second)
}

// Sum activities of day.
func sumActivitiesDay(day types.Day, now time.Time) (map[string]time.Duration, time.Duration) {
	var dur time.Duration
	var durs = make(map[string]time.Duration)

	for _, act := range day.Activities {
		d := act.End.Sub(act.Start)

		durs[act.Title] += d
		dur += d
	}

	if day.IsOccupied() {
		d := now.Sub(day.CurActivity.Start)
		durs[day.CurActivity.Title] += d
		dur += d
	}

	return durs, dur.Round(time.Second)
}

// Sum time spans.
func sumSpans(spans []types.Span) time.Duration {
	var dur time.Duration

	for _, span := range spans {
		dur += span.End.Sub(span.Start)
	}

	return dur.Round(time.Second)
}
