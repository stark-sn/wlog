package working

import (
	"time"

	"s-stark.net/code/wlog/types"
)

func startSpan(day types.Day, t time.Time) types.Day {
	if day.CurSpan == nil {
		span := types.Span{}
		span.Start = t
		day.CurSpan = &span
	}

	return day
}

func endBreak(day types.Day, t time.Time) types.Day {
	if day.CurBreak != nil {
		b := day.CurBreak
		day.CurBreak = nil
		b.End = t
		day.Breaks = append(day.Breaks, *b)
	}

	return day
}

func endActivity(day types.Day, t time.Time) types.Day {
	if day.CurActivity != nil {
		a := day.CurActivity
		day.CurActivity = nil
		a.End = t
		day.Activities = append(day.Activities, *a)
	}

	return day
}
