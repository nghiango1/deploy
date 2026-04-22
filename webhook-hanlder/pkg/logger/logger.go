package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Get() *slog.Logger {
	return logger
}

func SetupLogger() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger = slog.New(handler)
}
