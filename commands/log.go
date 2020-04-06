// This package provides the commands that can be executed by this application.
package commands

import (
	"log"
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

func logCommandFunc(args []string) {
	if len(args) != 2 {
		log.Fatal("Usage: log <activity> <duration>")
	}

	activity := args[0]
	dur, err := time.ParseDuration(args[1])

	if err != nil {
		log.Fatalf("Failed to parse activity duration, %v, %w", args[1], err)
	}

	if dur <= 0 {
		log.Fatalf("Can't log activity '%v' with zero/negative duragion '%v'.", activity, dur)
	}

	app.LogActivity(activity, dur)
}

