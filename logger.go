package docker_exposer

import (
	"log/slog"
	"os"
)

var logLevel = new(slog.LevelVar)
var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	logLevel.Set(slog.LevelDebug)
}

func DefaultLogger() *slog.Logger {
	return logger
}

func SetDefaultLogLevel(level slog.Level) {
	logLevel.Set(level)
}
