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

	dur, err := SumWorkingTimeDay(day, now)

	if err != nil {
		return err
	}

	activities, actDur, err := SumActivitiesDay(day, now)

	if err != nil {
		return err
	}

	fmt.Printf("%v %15v\n", date, dur.Truncate(time.Second))

	if len(activities) > 0 {
		fmt.Println(strings.Repeat("┄", 26))
		ReportActivities(activities)
		fmt.Println(strings.Repeat("┅", 26))
		fmt.Printf("%26v\n", actDur.Truncate(time.Second))
	}

	return nil
}

// Sum up work time of day.
func SumWorkingTimeDay(day types.Day, now time.Time) (time.Duration, error) {
	var dur time.Duration

	for i, span := range day.Spans {
		var end time.Time
		if span.End.IsZero() {
			if i != len(day.Spans) - 1 {
				return dur, fmt.Errorf("Unclosed span at %d", i)
			}
			end = now
		} else {
			end = span.End
		}

		dur += end.Sub(span.Start)
	}

	return dur, nil
}

// Sum activities of day.
func SumActivitiesDay(day types.Day, now time.Time) (map[string]time.Duration, time.Duration, error) {
	var dur time.Duration
	var durs = make(map[string]time.Duration)

	for i, act := range day.Activities {
		var end time.Time
		if act.End.IsZero() {
			if i != len(day.Activities) - 1 {
				return durs, dur, fmt.Errorf("Unclosed activity at %d", i)
			}
			end = now
		} else {
			end = act.End
		}

		d := end.Sub(act.Start)

		durs[act.Title] = durs[act.Title] + d
		dur += d
	}

	return durs, dur, nil
}
