// This package provides the commands that can be executed by this application.
package commands

import (
	"errors"
	"fmt"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(breakCommand)
}

var breakCommand = Cmd{
	Use: "break",
	Run: breakCommandFunc,
}

func breakCommandFunc(args []string) error {

	if len(args) != 1 {
		return errors.New("Usage: break <start|end>")
	}

	if args[0] == "start" {
		return app.StartBreak()
	} else if args[0] == "end" {
		return app.EndCurrentBreak()
	} else {
		return fmt.Errorf("Unknown break command '%v'.", args[0])
	}

	return nil
}

