package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"slices"
	"slogger/internal/commands"
	"slogger/internal/config"
	"slogger/pkg/foundation/app"
	"strings"
)

var args = os.Args

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

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
		getLogLevels(),
		config.GetConfig().GetLogDir(),
		config.GetConfig().GetLogKeepDays(),
	)
}

func getLogLevels() []slog.Level {
	logLevels := strings.Split(config.GetConfig().GetLogLevels(), ",")

	var slogLevels []slog.Level

	if slices.Index(logLevels, "any") == -1 {
		for _, logLevel := range logLevels {
			if logLevel == "" {
				continue
			}

			switch logLevel {
			case "debug":
				slogLevels = append(slogLevels, slog.LevelDebug)
			case "info":
				slogLevels = append(slogLevels, slog.LevelInfo)
			case "warn":
				slogLevels = append(slogLevels, slog.LevelWarn)
			case "error":
				slogLevels = append(slogLevels, slog.LevelError)
			default:
				panic(fmt.Errorf("unknown log level: %s", logLevel))
			}
		}
	}

	return slogLevels
}
