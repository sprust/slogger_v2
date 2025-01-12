package commands

import (
	"slogger/internal/commands/hello_command"
	foundationCommands "slogger/pkg/foundation/commands"
)

const (
	HelloCommandName = "hello"
)

var commands = map[string]foundationCommands.CommandInterface{
	HelloCommandName: &hello_command.Command{},
}

func GetCommands() map[string]foundationCommands.CommandInterface {
	return commands
}
