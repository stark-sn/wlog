// The working package provide means to log periods of working time.
package working

import (
	"errors"
	"fmt"
	"time"
	"s-stark.net/code/wlog/types"
)

// Start break.
func StartBreak(week types.Week, t time.Time) (types.Week, error) {
	if week.Days == nil {
		return week, errors.New("You're not logged in this week yet.")
	}

	date := types.Date(t)
	day, _ := week.Days[date]

	if day.IsOnBreak() {
		return week, fmt.Errorf("You're already taking a break.")
	}

	if day.IsOccupied() {
		return week, fmt.Errorf("You're currently occupied with %s.", day.CurActivity.Title)
	}

	if !day.IsIn() {
		return week, errors.New("You're currently not in.")
	}

	b := types.Span{Start: t}
	day.CurBreak = &b
	week.Days[date] = day

	return week, nil
}

// End current break.
func EndCurrentBreak(week types.Week, t time.Time) (types.Week, error) {
	if week.Days == nil {
		return week, errors.New("You're not logged in this week yet.")
	}

	date := types.Date(t)
	day, ok := week.Days[date]

	if !ok {
		return week, errors.New("You're not logged in today yet.")
	}

	if !day.IsOnBreak() {
		return week, errors.New("You are not taking a break right now.")
	}

	b := day.CurBreak
	b.End = t
	day.CurBreak = nil

	day.Breaks = append(day.Breaks, *b)

	week.Days[date] = day

	return week, nil
}

