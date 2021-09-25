package working

import (
	"testing"
	"time"

	"s-stark.net/code/wlog/types"
)

func TestStartActivity(t *testing.T) {
	week := types.Week{}

	week, err := StartActivity(week, "", time.Time{})

	if err == nil {
		t.Error("Able to start activity without comming in first")
	}

	week, _ = ComeIn(week, time.Time{})

	week, err = StartActivity(week, "", time.Time{})

	if err != nil {
		t.Error("Unable to start activity")
	}

	week, err = StartActivity(week, "", time.Time{})

	if err == nil {
		t.Error("Abtle to start activity twice")
	}

	week, _ = GoOut(week, time.Time{})

	week, err = StartActivity(week, "", time.Time{})

	if err == nil {
		t.Error("Abtle to start activity when out")
	}

	week, _ = ComeIn(week, time.Time{})
	week, _ = StartBreak(week, time.Time{})

	week, err = StartActivity(week, "", time.Time{})

	if err == nil {
		t.Error("Abtle to start activity while being on break")
	}
}

func TestEndCurrentActivity(t *testing.T) {
	week := types.Week{}

	week, err := EndCurrentActivity(week, time.Time{})

	if err == nil {
		t.Error("Able to end activity without being in all week")
	}

	week.Days = make(map[string]types.Day)

	week, err = EndCurrentActivity(week, time.Time{})

	if err == nil {
		t.Error("Able to end activity without being in")
	}

	week, _ = ComeIn(week, time.Time{})

	week, err = EndCurrentActivity(week, time.Time{})

	if err == nil {
		t.Error("Able to end activity without being occupied")
	}

	week, _ = StartActivity(week, "", time.Time{})
	week, err = EndCurrentActivity(week, time.Time{})

	if err != nil {
		t.Error("Unable to end activity")
	}
}

func TestLogActivity(t *testing.T) {
	week := types.Week{}

	week, err := LogActivity(week, "", time.Time{}, 0)

	if err == nil {
		t.Error("Able to log activity without being in all week")
	}

	week.Days = make(map[string]types.Day)

	week, err = LogActivity(week, "", time.Time{}, 0)

	if err == nil {
		t.Error("Able to log activity without being in")
	}

	week, _ = ComeIn(week, time.Date(2021, time.September, 26, 2, 0, 0, 0, time.UTC))

	week, err = LogActivity(week, "", time.Date(2021, time.September, 26, 0, 0, 0, 0, time.UTC), 0)

	if err == nil {
		t.Error("Able to log activity that starts before being in")
	}

	week, _ = StartActivity(week, "", time.Date(2021, time.September, 26, 2, 30, 0, 0, time.UTC))

	week, err = LogActivity(week, "", time.Date(2021, time.September, 26, 3, 0, 0, 0, time.UTC), 0)

	if err == nil {
		t.Error("Able to log activity while being occupied")
	}

	week, _ = EndCurrentActivity(week, time.Date(2021, time.September, 26, 2, 45, 0, 0, time.UTC))

	week, err = LogActivity(week, "", time.Date(2021, time.September, 26, 3, 0, 0, 0, time.UTC), 0)

	if err != nil {
		t.Error("Unable to log activity")
	}
}
