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

	var start time.Time

	if day.Spans != nil {
		start = day.Spans[0].Start
	} else {
		start = day.CurSpan.Start
	}

	fmt.Printf("You started at %s.\n", fmtTime(start))

	if day.IsOnBreak() {
		fmt.Printf("You are currently on a break since %s.\n", fmtTime(day.CurBreak.Start))
	} else if day.IsOccupied() {
		fmt.Printf("You are currently working on '%s' since %s.\n", day.CurActivity.Title, fmtTime(day.CurActivity.Start))
	} else if !day.IsIn() {
		fmt.Printf("You are currently not in, you left at %s.\n", fmtTime(day.Spans[len(day.Spans)-1].End))
	} else {
		fmt.Println("You are currently slacking off.")
	}

	spanTime := sumWorkingTimeDay(day, t)
	breakTime := sumBreakTime(day, t)

	workTime := spanTime - breakTime

	fmt.Printf("You were on break for %s.\n", fmtDuration(breakTime))
	fmt.Printf("You worked for %s.\n", fmtDuration(workTime))

	return nil
}
