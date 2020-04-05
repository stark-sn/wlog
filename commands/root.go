// Package commands provides means to register and execute commands.
package commands

import (
	"fmt"
	"log"
	"os"
)

// Command definition.
type Cmd struct {
	Use string
	Run func(args []string) error
}

var commands map[string]Cmd = make(map[string]Cmd)

var gotDefault bool
var defaultCommand Cmd
var defaultArgs []string

// Add a command.
func AddCommand(command Cmd) {
	if command.Run == nil {
		log.Fatal("Trying to add a command without a handler function.")
	}

	commands[command.Use] = command
}

// Add a command as the applications default command.
func AddDefaultCommand(command Cmd, args []string) {
	if gotDefault {
		log.Fatal("Default command is already defined.")
	}

	gotDefault = true
	defaultCommand = command
	defaultArgs = args

	AddCommand(command)
}

// Execute uses os.Args to determine what command should be executed.
func Execute() error {

	var args []string
	var command Cmd

	binName := os.Args[0]

	if len(os.Args) < 2 {
		if !gotDefault {
			return fmt.Errorf("Usage: %s", binName)
		}

		command = defaultCommand
		args = defaultArgs
	} else {
		cmdName := os.Args[1]
		args = os.Args[2:]

		c, exists := commands[cmdName]

		if !exists {
			return fmt.Errorf("Unknown command '%s'", cmdName)
		}

		command = c
	}

	return command.Run(args)
}
