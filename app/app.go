package app

import (
	"fmt"
	"time"
	"os/user"
	"path"
	"s-stark.net/code/wlog/types"
	"s-stark.net/code/wlog/reporting"
	"s-stark.net/code/wlog/persistence"
	"s-stark.net/code/wlog/working"
)

const dir = ".wlog"

var now time.Time = time.Now()
var weekFile string
var week types.Week

func In() error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	week, err = working.ComeIn(week, now)

	if err != nil {
		return err
	}

	err = writeWeekFile()
	if err != nil {
		return err
	}

	return nil
}

func Out() error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	week, err = working.GoOut(week, now)

	if err != nil {
		return err
	}

	err = writeWeekFile()
	if err != nil {
		return err
	}

	return nil
}

func StartBreak() error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	week, err = working.StartBreak(week, now)

	if err != nil {
		return err
	}

	err = writeWeekFile()
	if err != nil {
		return err
	}

	return nil
}

func EndCurrentBreak() error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	week, err = working.EndCurrentBreak(week, now)

	if err != nil {
		return err
	}

	err = writeWeekFile()
	if err != nil {
		return err
	}

	return nil
}

func ReportDay() error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	return reporting.ReportDay(week, now)
}

func ReportWeek() error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	return reporting.ReportWeek(week, now)
}

func StartActivity(task string) error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	week, err = working.StartActivity(week, task, now)

	if err != nil {
		return err
	}

	err = writeWeekFile()
	if err != nil {
		return err
	}

	return nil
}

func EndCurrentActivity() error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	week, err = working.EndCurrentActivity(week, now)

	if err != nil {
		return err
	}

	err = writeWeekFile()
	if err != nil {
		return err
	}

	return nil
}

func LogActivity(task string, duration time.Duration) error {
	err := readWeekFile()
	if err != nil {
		return err
	}

	week, err = working.LogActivity(week, task, now, duration)

	if err != nil {
		return err
	}

	err = writeWeekFile()
	if err != nil {
		return err
	}

	return nil
}

// helper functions

// Read week data from file
func readWeekFile() error {
	user, err := user.Current()

	if err != nil {
		return fmt.Errorf("Failed to get current user, %w", err)
	}

	y, w := now.ISOWeek()
	weekFile = path.Join(user.HomeDir, dir, fmt.Sprintf("%d", y), fmt.Sprintf("%02d", w)) + ".json"

	week, err = persistence.Read(weekFile)

	if err != nil {
		return err
	}

	return nil
}

// Write week data to file
func writeWeekFile() error {
	return persistence.Write(weekFile, week)
}

