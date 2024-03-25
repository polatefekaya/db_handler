package log

import "log/slog"

var logLevel = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

type Options struct {
	Log *slog.HandlerOptions
}

func newHandlerOptions(level string) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel[level],
	}
}
