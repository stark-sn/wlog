// The reporting package provide means to create work time reports.
package reporting

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
	"s-stark.net/code/wlog/types"
)

// Create work time report for week.
func ReportWeek(week types.Week, now time.Time) error {
	if week.Days == nil {
		return errors.New("You're not in this week?")
	}

	var dates []string
	for date, _ := range week.Days {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	var dur time.Duration

	for _, date := range dates {
		day, _ := week.Days[date]
		dayDur, err := SumWorkingTimeDay(day, now)
		if err != nil {
			return err
		}

		activities, actDur, err := SumActivitiesDay(day, now)
		if err != nil {
			return err
		}

		fmt.Printf("%v %15v\n", date, dayDur.Truncate(time.Second))

		if len(activities) > 0 {
			fmt.Println(strings.Repeat("┄", 26))
			ReportActivities(activities)
			fmt.Println(strings.Repeat("┅", 26))
			fmt.Printf("%26s\n", actDur.Truncate(time.Second))
			fmt.Println(strings.Repeat("─", 26))
			fmt.Println()
		}

		dur += dayDur
	}

	fmt.Println(strings.Repeat("═", 26))
	fmt.Printf("Week %21v\n",dur.Truncate(time.Second))
	return nil
}

