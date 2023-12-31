// The working package provide means to log periods of working time.
package working

import (
	"errors"
	"fmt"
	"s-stark.net/code/wlog/types"
	"time"
)

// Come in for work.
func ComeIn(week types.Week, t time.Time) (types.Week, error) {
	// New week
	if week.Days == nil {
		week.Days = make(map[string]types.Day)
	}

	date := types.Date(t)
	day, _ := week.Days[date]

	if day.IsIn() {
		return week, fmt.Errorf("You're already in since %v.", day.CurSpan.Start)
	}

	// Start new work span
	day = startSpan(day, t)

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

	if !day.IsIn() {
		return week, errors.New("You're are currently out.")
	}

	// End current span
	span := day.CurSpan
	day.CurSpan = nil
	span.End = t
	day.Spans = append(day.Spans, *span)

	// End current activity and break if applicable
	day = endBreak(day, t)
	day = endActivity(day, t)

	week.Days[date] = day

	return week, nil
}
