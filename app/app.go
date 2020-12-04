package app

import (
	"fmt"
	"log"
	"os/user"
	"path"
	"s-stark.net/code/wlog/persistence"
	"s-stark.net/code/wlog/reporting"
	"s-stark.net/code/wlog/types"
	"s-stark.net/code/wlog/working"
	"time"
)

const dir = ".wlog"

var now time.Time = time.Now()
var weekFile string
var week types.Week
var err error

func init() {
	readWeekFile()
}

func In() {
	save(working.ComeIn(week, now))
}

func Out() {
	save(working.GoOut(week, now))
}

func StartBreak() {
	save(working.StartBreak(week, now))
}

func EndCurrentBreak() {
	save(working.EndCurrentBreak(week, now))
}

func Status() {
	reporting.Status(week, now)
}

func Timesheet() {
	reporting.Timesheet(week, now)
}

func ReportDay() {
	reporting.ReportDay(week, now)
}

func ReportWeek() {
	reporting.ReportWeek(week, now)
}

func StartActivity(task string) {
	save(working.StartActivity(week, task, now))
}

func EndCurrentActivity() {
	save(working.EndCurrentActivity(week, now))
}

func LogActivity(task string, duration time.Duration) {
	save(working.LogActivity(week, task, now, duration))
}

// helper functions

// Read week data from file
func readWeekFile() {
	user, err := user.Current()

	if err != nil {
		log.Fatalf("Failed to get current user, %v", err)
	}

	y, w := now.ISOWeek()
	weekFile = path.Join(user.HomeDir, dir, fmt.Sprintf("%d", y), fmt.Sprintf("%02d", w)) + ".json"

	week, err = persistence.Read(weekFile)

	if err != nil {
		log.Fatalf("Failed to read from %v, %v", weekFile, err)
	}

}

// save changes if no error occoured
func save(w types.Week, e error) {
	if e != nil {
		log.Fatal(e)
	}

	week = w
	writeWeekFile()
}

// Write week data to file
func writeWeekFile() {
	err := persistence.Write(weekFile, week)
	if err != nil {
		log.Fatalf("Failed to write to file %v, %v", weekFile, err)
	}
}
