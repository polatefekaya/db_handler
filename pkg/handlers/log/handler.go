package log

import (
	"io"
	"log/slog"
)

type Handler struct {
	Handler slog.Handler
}

func newHandler(handlerType string, w io.Writer, opts *slog.HandlerOptions) *slog.Handler {
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
		return nil
	}

	return &h
}
