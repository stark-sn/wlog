package working

import (
	"testing"
	"time"

	"s-stark.net/code/wlog/types"
)

func TestLogNonWorkingTime(t *testing.T) {
	week := types.Week{}
	now := time.Time{}

	week, err := LogNonWorkingTime(week, now, "", 0)
	if err != nil || len(week.Days[types.Date(now)].NonWorkingTime) != 1 {
		t.Error("Unable to log non working time")
	}

	week, err = LogNonWorkingTime(week, now, "", 0)
	if err != nil || len(week.Days[types.Date(now)].NonWorkingTime) != 2 {
		t.Error("Unable to log non working time twice")
	}

}
