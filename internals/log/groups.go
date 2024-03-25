package log

import (
	"log/slog"
	"math/rand"
)

type Group struct {
	Log *slog.Logger
}

// TODO: Not efficient, like shit
func WithUsers() *CustomLog {
	userGroup := slog.Group(
		"users",
		slog.String("name", "polat"),
		slog.Int("id", rand.Intn(1000)),
	)
	var cl = NewSLog()
	cl.Log = initLog.Log.With(userGroup)
	return cl
}
