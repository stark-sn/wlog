// This package provides the commands that can be executed by this application.
package commands

import (
	"log"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(startCommand)
}

var startCommand = Cmd{
	Use: "start",
	Run: startCommandFunc,
}

func startCommandFunc(args []string) {
	if len(args) != 1 {
		log.Fatal("Usage: start <activity>")
	}

	app.StartActivity(args[0])
}
