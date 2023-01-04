package working

import (
	"testing"
	"time"

	"s-stark.net/code/wlog/types"
)

func TestStartBreak(t *testing.T) {
	week := types.Week{}
	week, err := StartBreak(week, time.Time{})

	if err != nil {
		t.Error("Unable to start break")
	}

	week, err = StartBreak(week, time.Time{})

	if err == nil {
		t.Error("Abtle to start break twice")
	}
}

func TestEndCurrentBreak(t *testing.T) {
	week := types.Week{}

	week, err := EndCurrentBreak(week, time.Time{})

	if err == nil {
		t.Error("Able to end break without being in all week")
	}

	week.Days = make(map[string]types.Day)

	week, err = EndCurrentBreak(week, time.Time{})

	if err == nil {
		t.Error("Able to end break without being in")
	}

	week, _ = ComeIn(week, time.Time{})

	week, err = EndCurrentBreak(week, time.Time{})

	if err == nil {
		t.Error("Able to end break without being on break")
	}

	week, _ = StartBreak(week, time.Time{})
	week, err = EndCurrentBreak(week, time.Time{})

	if err != nil {
		t.Error("Unable to end break")
	}
}
