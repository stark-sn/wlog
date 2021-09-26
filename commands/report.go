// This package provides the commands that can be executed by this application.
package commands

import (
	"log"

	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(timesheetCommand)
	AddDefaultCommand(reportCommand, []string{"day"})
}

var timesheetCommand = Cmd{
	Use: "timesheet",
	Run: timesheetCommandFunc,
}

var reportCommand = Cmd{
	Use: "report",
	Run: reportCommandFunc,
}

func timesheetCommandFunc(args []string) {
	app.Timesheet()
}

func reportCommandFunc(args []string) {
	if len(args) != 1 {
		log.Fatal("Supported arguments are 'day' and 'week'")
	}

	switch reportType := args[0]; reportType {
	case "day":
		app.ReportDay()
	case "week":
		app.ReportWeek()
	default:
		log.Fatalf("Unsupported report type '%v'", reportType)
	}
}
