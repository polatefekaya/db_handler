package usecases

import (
	"errors"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"os"
)

var logLevel = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

var logHandler = map[string]slog.Handler{
	"json": slog.NewJSONHandler(nil, nil),
	"text": slog.NewTextHandler(nil, nil),
}

func NewCustomLogger() *slog.Logger {
	l := createLogger("json", "debug", os.Stderr)

	return l
}

func (l *) WithUsers(){
	
}

func createLogger(hType, hLevel string, w io.Writer) *slog.Logger {
	h, err := newHandler(hType, w, newHandlerOptions(hLevel))
	if err != nil {
		log.Panic(err)
	}
	return slog.New(*h)
}

func newHandler(handlerType string, w io.Writer, opts *slog.HandlerOptions) (*slog.Handler, error) {
	var h slog.Handler

	switch handlerType {
	case "text":
		h = slog.NewTextHandler(w, opts)
	case "json":
		h = slog.NewJSONHandler(w, opts)
	default:
		h = nil
	}

	if h == nil {
		return nil, errors.New("Handler can not be created, nil")
	}

	return &h, nil
}

func newHandlerOptions(level string) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel[level],
	}
}

func logtest() {

	userGroup := slog.Group(
		"users",
		slog.String("name", "polat"),
		slog.Int("id", rand.Intn(1000)),
	)

	reqGroup := slog.Group(
		"requests",
	)

	//slog.SetDefault(logger)
	slog.Debug("Debug")
	slog.Info(
		"Info", userGroup, reqGroup,
	)
	slog.Warn("Warn")
	slog.Error("Error")
}
