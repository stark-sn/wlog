// wlog is console application that allows you to track working time and activities.
package main

import (
	"log"
	"s-stark.net/code/wlog/commands"
)

func init() {
	log.SetFlags(0)
}

func main() {
	commands.Execute()
}
