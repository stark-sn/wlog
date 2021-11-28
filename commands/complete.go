// This package provides the commands that can be executed by this application.
package commands

import (
	"log"

	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(findCommand)
}

var findCommand = Cmd{
	Use: "find",
	Run: findCommandFunc,
}

func findCommandFunc(args []string) {
	if len(args) != 1 {
		log.Fatal("Usage: find <text>")
	}

	app.FindActivity(args[0])
}
