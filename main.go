// wlog is console application that allows you to track working time and activities.
package main

import (
	"fmt"
	"os"
	"s-stark.net/code/wlog/commands"
)

func main() {
	err := commands.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

