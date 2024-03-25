package config

import (
	"DatabaseHandler/pkg/data/models/Players"
	"DatabaseHandler/pkg/handlers/log"
	"DatabaseHandler/pkg/usecases"
)

type AppConfig struct {
	Conn       string
	Football   *usecases.FootballUsecase
	Automation *usecases.AutomationUseCase
	Logger     *log.CustomLog
}

type AppCache struct {
	PlayerModel *Players.PlayerRoot
}

type Logger struct {
	Log *log.CustomLog
}

func NewLogger() *Logger {
	return &Logger{Log: log.NewSLog()}
}
