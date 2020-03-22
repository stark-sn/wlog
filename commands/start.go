// This package provides the commands that can be executed by this application.
package commands

import (
	"errors"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(startCommand)
}

var startCommand = Cmd{
	Use: "start",
	Run: startCommandFunc,
}

func startCommandFunc(args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: start <activity>")
	}

	return app.StartActivity(args[0])
}

