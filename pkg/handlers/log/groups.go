package log

import (
	"log/slog"
	"math/rand"
)

type Group struct {
	Log *slog.Logger
}

func (l *CustomLog) WithUsers() *CustomLog {
	userGroup := slog.Group(
		"users",
		slog.String("name", "polat"),
		slog.Int("id", rand.Intn(1000)),
	)
	l.Log = l.Log.With(userGroup)
	return l
}
