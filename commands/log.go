// This package provides the commands that can be executed by this application.
package commands

import (
	"errors"
	"fmt"
	"time"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(logCommand)
}

var logCommand = Cmd{
	Use: "log",
	Run: logCommandFunc,
}

func logCommandFunc(args []string) error {
	if len(args) != 2 {
		return errors.New("Usage: log <activity> <duration>")
	}

	activity := args[0]
	dur, err := time.ParseDuration(args[1])

	if err != nil {
		return err
	}

	if dur <= 0 {
		return fmt.Errorf("Can't log activity '%v' with zero/negative duragion '%v'.", activity, dur)
	}

	return app.LogActivity(activity, dur)
}

