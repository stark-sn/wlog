// The working package provide means to log periods of working time.
package working

import (
	"errors"
	"fmt"
	"time"
	"s-stark.net/code/wlog/types"
)

// Come in for work.
func ComeIn(week types.Week, t time.Time) (types.Week, error) {
	// New week
	if week.Days == nil {
		week.Days = make(map[string]types.Day)
	}

	date := types.Date(t)
	day, _ := week.Days[date]

	span, wasInOnce := day.CurrentSpan()

	// Current span has not ended yet
	if wasInOnce && !span.Completed() {
		return week, fmt.Errorf("You're already in since %v.", span.Start)
	}

	// Start new work span
	span = types.Span{}
	span.Start = t
	day.Spans = append(day.Spans, span)

	week.Days[date] = day

	return week, nil
}

// Leave work.
func GoOut(week types.Week, t time.Time) (types.Week, error) {
	if week.Days == nil {
		return week, errors.New("You're not logged in this week. Did you over midnight on a sunday?")
	}

	date := types.Date(t)
	day, wasInToday := week.Days[date]

	if !wasInToday {
		return week, errors.New("You're not logged in today. Did you work over midnight?")
	}


	span, wasInOnce := day.CurrentSpan()

	if !wasInOnce || span.Completed() {
		return week, errors.New("You're are currently out.")
	}

	// End currently running activity
	if len(day.Activities) > 0 {
		j := len(day.Activities) - 1

		if day.Activities[j].End.IsZero() {
			day.Activities[j].End = t
		}
	}

	span.End = t;
	day.Spans[len(day.Spans) - 1] = span
	week.Days[date] = day

	return week, nil
}

