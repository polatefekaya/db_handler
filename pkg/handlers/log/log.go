package log

import (
	"io"
	"log"
	"log/slog"
	"os"
)

type CustomLog struct {
	Log *slog.Logger
}

func newCustomLogger() *slog.Logger {
	l := createLogger("json", "debug", os.Stderr)

	return l
}
func NewSLog() *CustomLog {
	return &CustomLog{
		Log: newCustomLogger(),
	}
}

func createLogger(hType, hLevel string, w io.Writer) *slog.Logger {
	h := newHandler(hType, w, newHandlerOptions(hLevel))
	if h == nil {
		log.Panic("Handler is nil!")
	}
	return slog.New(*h)
}
