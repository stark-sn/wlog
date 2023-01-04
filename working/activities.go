// The working package provide means to log periods of working time.
package working

import (
	"errors"
	"fmt"
	"s-stark.net/code/wlog/types"
	"time"
)

// Start new activity.
func StartActivity(week types.Week, activity string, t time.Time) (types.Week, error) {
	if week.Days == nil {
		week.Days = make(map[string]types.Day)
	}

	date := types.Date(t)
	day, _ := week.Days[date]

	if day.IsOccupied() {
		if day.CurActivity.Title == activity {
			return week, errors.New("You are already occupied with this activity.")
		}

		day = endActivity(day, t)
	}

	if day.IsOnBreak() {
		day = endBreak(day, t)
	}

	if !day.IsIn() {
		day = startSpan(day, t)
	}

	act := types.Activity{Title: activity}
	act.Start = t
	day.CurActivity = &act
	week.Days[date] = day

	return week, nil
}

// End activity that is currently ongoing.
func EndCurrentActivity(week types.Week, t time.Time) (types.Week, error) {
	if week.Days == nil {
		return week, errors.New("You're not logged in this week yet.")
	}

	date := types.Date(t)
	day, ok := week.Days[date]

	if !ok {
		return week, errors.New("You're not logged in today yet.")
	}

	if !day.IsOccupied() {
		return week, errors.New("You are not occupied with any activity at the moment.")
	}

	act := day.CurActivity
	act.End = t
	day.CurActivity = nil

	day.Activities = append(day.Activities, *act)

	week.Days[date] = day

	return week, nil
}

// Log an activity in the past.
func LogActivity(week types.Week, activity string, t time.Time, dur time.Duration) (types.Week, error) {
	if week.Days == nil {
		return week, errors.New("You're not logged in this week yet.")
	}

	date := types.Date(t)
	day, _ := week.Days[date]

	act := types.Activity{Title: activity}
	act.Start = t.Add(-1 * dur)
	act.End = t

	if !day.IsIn() {
		return week, errors.New("You're currently not in.")
	}

	if act.Start.Before(day.CurSpan.Start) {
		return week, errors.New("Can't log activity that starts before you were in.")
	}

	if day.IsOccupied() {
		return week, fmt.Errorf("You are still occupied with this activity '%v'", day.CurActivity.Title)
	}

	day.Activities = append(day.Activities, act)
	week.Days[date] = day

	return week, nil
}
