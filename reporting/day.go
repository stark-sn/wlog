// The reporting package provide means to create work time reports.
package reporting

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"s-stark.net/code/wlog/types"
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

	dur := SumWorkingTimeDay(day, now)
	breakTime := SumBreakTime(day, now)
	activities, sumActs := SumActivitiesDay(day, now)

	dur -= breakTime

	fmt.Printf("%v %15v\n", date, dur.Truncate(time.Second))
	fmt.Println(strings.Repeat("┄", 26))
	fmt.Printf("Break % 20v\n", breakTime.Truncate(time.Second))

	if len(activities) > 0 {
		fmt.Println(strings.Repeat("┄", 26))
		ReportActivities(activities)
		fmt.Println(strings.Repeat("┅", 26))
		fmt.Printf("%26v\n", sumActs.Truncate(time.Second))
	}

	return nil
}

// Sum up work time of day.
func SumWorkingTimeDay(day types.Day, now time.Time) time.Duration {
	dur := sumSpans(day.Spans)

	if day.IsIn() {
		dur += now.Sub(day.CurSpan.Start)
	}

	return dur
}

// Sum up break tiem of day.
func SumBreakTime(day types.Day, now time.Time) time.Duration {
	dur := sumSpans(day.Breaks)

	if day.IsOnBreak() {
		dur += now.Sub(day.CurBreak.Start)
	}

	return dur
}

// Sum activities of day.
func SumActivitiesDay(day types.Day, now time.Time) (map[string]time.Duration, time.Duration) {
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

	return durs, dur
}

// Sum time spans.
func sumSpans(spans []types.Span) time.Duration {
	var dur time.Duration

	for _, span := range spans {
		dur += span.End.Sub(span.Start)
	}

	return dur
}
