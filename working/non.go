package working

import (
	"time"

	"s-stark.net/code/wlog/types"
)

func LogNonWorkingTime(week types.Week, t time.Time, title string, duration time.Duration) (types.Week, error) {
	if week.Days == nil {
		week.Days = make(map[string]types.Day)
	}

	date := types.Date(t)
	day, _ := week.Days[date]

	day.NonWorkingTime = append(day.NonWorkingTime, types.NonWorkingTime{
		title,
		duration,
	})

	week.Days[date] = day

	return week, nil
}
