package commands

import (
	"slogger/internal/commands/hello_command"
	"slogger/internal/commands/serve_grpc_command"
	foundationCommands "slogger/pkg/foundation/commands"
)

const (
	HelloCommandName     = "hello"
	ServeGrpcCommandName = "serve:grpc"
)

var commands = map[string]foundationCommands.CommandInterface{
	HelloCommandName:     &hello_command.Command{},
	ServeGrpcCommandName: &serve_grpc_command.Command{},
}

func GetCommands() map[string]foundationCommands.CommandInterface {
	return commands
}
