package working

import (
	"s-stark.net/code/wlog/types"
	"testing"
	"time"
)

var (
	t1 time.Time = time.Date(2020, time.December, 3, 6, 0, 0, 0, time.UTC)
	t2 time.Time = time.Date(2020, time.December, 2, 6, 0, 0, 0, time.UTC)
	t3 time.Time = time.Date(2020, time.December, 1, 6, 0, 0, 0, time.UTC)
)

func TestComeIn(t *testing.T) {
	week := types.Week{}

	week, err := ComeIn(week, t1)

	if err != nil {
		t.Errorf("Unable to ComeIn to new week")
	}

	week, err = ComeIn(week, t1)

	if err == nil {
		t.Errorf("Able to come into week twice")
	}
}

func TestGoOut(t *testing.T) {
	week := types.Week{}

	week, err := GoOut(week, t1)

	if err == nil {
		t.Errorf("Able to get out of week without being in all week")
	}

	week, err = ComeIn(week, t1)

	week, err = GoOut(week, t2)

	if err == nil {
		t.Errorf("Able to get out ather midnight")
	}

	week, err = GoOut(week, t1)
	week, err = GoOut(week, t1)

	if err == nil {
		t.Errorf("Able to get out twice")
	}

	week, err = ComeIn(week, t3)
	week, err = StartBreak(week, t3)
	week, err = GoOut(week, t3)
	day, _ := week.Days["2020-12-01"]

	if len(day.Spans) != 1 || len(day.Breaks) != 1 {
		t.Errorf("Span & Break not closed")
	}

	week, err = ComeIn(week, t3)
	week, err = StartActivity(week, "sht", t3)
	week, err = GoOut(week, t3)
	day, _ = week.Days["2020-12-01"]

	if len(day.Spans) != 2 || len(day.Activities) != 1 {
		t.Errorf("Span & Activity not closed")
	}
}
