// This package provides the commands that can be executed by this application.
package commands

import (
	"log"
	"time"

	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(nonCommand)
}

var nonCommand = Cmd{
	Use: "non",
	Run: nonCommandFunc,
}

func nonCommandFunc(args []string) {
	if len(args) != 2 {
		log.Fatal("Usage: non <title> <duration>")
	}

	title := args[0]
	dur, err := time.ParseDuration(args[1])

	if err != nil {
		log.Fatalf("Failed to parse non working time duration, %v, %v", args[1], err)
	}

	if dur <= 0 {
		log.Fatalf("Can't log non working time '%v' with zero/negative duragion '%v'.", title, dur)
	}

	app.LogNonWorkingTime(title, dur)
}
