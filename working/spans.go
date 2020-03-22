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

	// Check if last span is finished
	if len(day.Spans) > 0 {
		span := day.Spans[len(day.Spans) - 1]
		if span.End.IsZero() {
			return week, fmt.Errorf("You're already in since %v.", span.Start)
		}
	}

	day.Spans = append(day.Spans, types.Span{Start: t})

	week.Days[date] = day

	return week, nil
}

// Leave work.
func GoOut(week types.Week, t time.Time) (types.Week, error) {
	if week.Days == nil {
		return week, errors.New("You're not logged in this week. Did you over midnight on a sunday?")
	}

	date := types.Date(t)
	day, ok := week.Days[date]

	if !ok {
		return week, errors.New("You're not logged in today. Did you work over midnight?")
	}

	if day.Spans == nil || len(day.Spans) == 0 {
		return week, errors.New("This should not happen. Did you mess around with the data file?")
	}

	i := len(day.Spans) - 1
	if day.Spans[i].Start.IsZero() {
		return week, errors.New("This can not happen. You did mess around with the data file!")
	}

	if !day.Spans[i].End.IsZero() {
		return week, errors.New("You're are currently out.")
	}

	// End currently running activity
	if len(day.Activities) > 0 {
		j := len(day.Activities) - 1

		if day.Activities[j].End.IsZero() {
			day.Activities[j].End = t
		}
	}

	day.Spans[i].End = t;
	week.Days[date] = day

	return week, nil
}

