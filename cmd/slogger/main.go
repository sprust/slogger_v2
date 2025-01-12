package main

import (
	"log/slog"
	"os"
	"slogger/internal/commands"
	"slogger/pkg/foundation/app"
)

var args = os.Args

func main() {
	var commandName string
	var commandArgs []string

	if len(args) > 1 {
		commandName = args[1]
	}

	if len(args) > 2 {
		commandArgs = args[2:]
	}

	newApp := app.NewApp(
		commands.GetCommands(),
		initAppConfig(),
	)

	newApp.Start(commandName, commandArgs)
}

func initAppConfig() *app.Config {
	return app.NewConfig(
		[]slog.Level{},
		"storage/logs",
		3,
	)
}
