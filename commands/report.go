// This package provides the commands that can be executed by this application.
package commands

import (
	"errors"
	"fmt"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddDefaultCommand(reportCommand, []string{"day"})
}

var reportCommand = Cmd{
	Use: "report",
	Run: reportCommandFunc,
}

func reportCommandFunc(args []string) error {
	if len(args) != 1 {
		return errors.New("Supported arguments are 'day' and 'week'")
	}

	switch reportType := args[0]; reportType {
		case "day":
			return app.ReportDay()
		case "week":
			return app.ReportWeek()
		default:
			return fmt.Errorf("Unsupported report type '%v'", reportType)
	}
}

