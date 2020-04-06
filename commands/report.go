// This package provides the commands that can be executed by this application.
package commands

import (
	"log"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddDefaultCommand(reportCommand, []string{"day"})
}

var reportCommand = Cmd{
	Use: "report",
	Run: reportCommandFunc,
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

