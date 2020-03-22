// This package provides the commands that can be executed by this application.
package commands

import (
	"errors"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(endCommand)
}

var endCommand = Cmd{
	Use: "end",
	Run: endCommandFunc,
}

func endCommandFunc(args []string) error {
	if len(args) > 0 {
		return errors.New("End does not support any arguments.")
	}

	return app.EndCurrentActivity()
}

