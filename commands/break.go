// This package provides the commands that can be executed by this application.
package commands

import (
	"log"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(breakCommand)
}

var breakCommand = Cmd{
	Use: "break",
	Run: breakCommandFunc,
}

func breakCommandFunc(args []string) {

	if len(args) != 1 {
		log.Fatal("Usage: break <start|end>")
	}

	if args[0] == "start" {
		app.StartBreak()
	} else if args[0] == "end" {
		app.EndCurrentBreak()
	} else {
		log.Fatalf("Unknown break command '%v'.", args[0])
	}
}
