// This package provides the commands that can be executed by this application.
package commands

import (
	"log"
	"s-stark.net/code/wlog/app"
)

func init() {
	AddCommand(endCommand)
}

var endCommand = Cmd{
	Use: "end",
	Run: endCommandFunc,
}

func endCommandFunc(args []string) {
	if len(args) > 0 {
		log.Fatal("End does not support any arguments.")
	}

	app.EndCurrentActivity()
}
