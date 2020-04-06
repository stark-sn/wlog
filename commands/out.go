// This package provides the commands that can be executed by this application.
package commands

import (
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(outCommand)
}

var outCommand = Cmd{
	Use: "out",
	Run: outCommandFunc,
}

func outCommandFunc(args []string) {
	app.Out()
}

