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

var initLog *CustomLog

func newCustomLogger() *slog.Logger {
	l := createLogger("json", "info", os.Stderr)

	return l
}
func NewSLog() *CustomLog {
	initLog = &CustomLog{Log: newCustomLogger()}
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

func INFO(s string, a ...any) {
	initLog.Log.Info(s, a...)
}
func DEBUG(s string, a ...any) {
	initLog.Log.Debug(s, a...)
}
func WARN(s string, a ...any) {
	initLog.Log.Warn(s, a...)
}
func (l *CustomLog) INFO(s string, a ...any) {
	l.Log.Info(s, a...)
}
func (l *CustomLog) DEBUG(s string, a ...any) {
	l.Log.Debug(s, a...)
}
func (l *CustomLog) WARN(s string, a ...any) {
	l.Log.Warn(s, a...)
}
