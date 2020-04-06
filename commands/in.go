// This package provides the commands that can be executed by this application.
package commands

import (
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(inCommand)
}

var inCommand = Cmd{
	Use: "in",
	Run: inCommandFunc,
}

func inCommandFunc(args []string) {
	app.In()
}

