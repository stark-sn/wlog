// The working package provide means to log periods of working time.
package working

import (
	"errors"
	"fmt"
	"time"
	"s-stark.net/code/wlog/types"
)

// Start new activity.
func StartActivity(week types.Week, activity string, t time.Time) (types.Week, error) {
	if week.Days == nil {
		return week, errors.New("You're not logged in this week yet.")
	}

	date := types.Date(t)
	day, _ := week.Days[date]

	// Check if last activity is finished
	if len(day.Activities) > 0 {
		activity := day.Activities[len(day.Activities) - 1]
		if activity.End.IsZero() {
			return week, fmt.Errorf("You're currently occupied with this activity '%v'.", activity.Title)
		}
	}

	if len(day.Spans) > 0 {
		i := len(day.Spans) - 1
		if !day.Spans[i].End.IsZero() {
			return week, errors.New("You're currently not in.")
		}
	} else {
		return week, errors.New("You're not in today.")
	}

	act := types.Activity{Title: activity}
	act.Start = t
	day.Activities = append(day.Activities, act)
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

	if day.Activities == nil || len(day.Activities) == 0 {
		return week, errors.New("You did not start any activitys today.")
	}

	i := len(day.Activities) - 1
	if day.Activities[i].Start.IsZero() {
		return week, errors.New("This can not happen. You did mess around with the data file!")
	}

	if !day.Activities[i].End.IsZero() {
		return week, errors.New("You are not occupied with any activity at the moment.")
	}

	day.Activities[i].End = t;
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

	if len(day.Spans) > 0 {
		i := len(day.Spans) - 1
		if !day.Spans[i].End.IsZero() {
			return week, errors.New("You're currently not in.")
		}

		if act.Start.Before(day.Spans[i].Start) {
			return week, errors.New("Can't log activity that starts before you were in.")
		}
	} else {
		return week, errors.New("You're not in today.")
	}

	// Check if last activity was finished
	if len(day.Activities) > 0 {
		act := day.Activities[len(day.Activities) - 1]
		if act.End.IsZero() {
			return week, fmt.Errorf("You are still occupied with this activity '%v'", act.Title)
		}
	}

	day.Activities = append(day.Activities, act)
	week.Days[date] = day

	return week, nil
}

