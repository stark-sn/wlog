// The reporting package provide means to create work time reports.
package reporting

import (
	"fmt"
	"s-stark.net/code/wlog/types"
	"time"
)

// Display current work status.
func Status(week types.Week, t time.Time) error {
	if week.Days == nil {
		fmt.Println("You're were not in this week.")
		return nil
	}

	date := types.Date(t)
	day, in := week.Days[date]

	if !in {
		fmt.Println("You're were not in today.")
		return nil
	}

	if day.IsOnBreak() {
		fmt.Printf("You are currently on a break since %s.\n", fmtTime(day.CurBreak.Start))
	} else if day.IsOccupied() {
		fmt.Printf("You are currently working on '%s' since %s.\n", day.CurActivity.Title, fmtTime(day.CurActivity.Start))
	} else if !day.IsIn() {
		if len(day.Spans) == 0 {
			fmt.Println("You never logged in today")
		} else {
			fmt.Printf("You are currently not in, you left at %s.\n", fmtTime(day.Spans[len(day.Spans)-1].End))
		}
	} else {
		fmt.Println("You are currently not working on any task.")
	}

	return nil
}
