package log

import (
	"log/slog"
	"strings"
)

var logLevel = map[string]slog.Level{
	"debug":     slog.LevelDebug,
	"deepDebug": slog.LevelDebug,
	"info":      slog.LevelInfo,
	"deepInfo":  slog.LevelInfo,
	"warn":      slog.LevelWarn,
	"deepWarn":  slog.LevelWarn,
	"error":     slog.LevelError,
	"deepError": slog.LevelError,
}

type Options struct {
	Log *slog.HandlerOptions
}

func newHandlerOptions(level string) *slog.HandlerOptions {
	var as = strings.Contains(strings.ToLower(level), "deep")
	return &slog.HandlerOptions{
		AddSource: as,
		Level:     logLevel[level],
	}
}
