package config

import (
	"DatabaseHandler/internals/log"
	"DatabaseHandler/pkg/data/models/Players"
	"DatabaseHandler/pkg/usecases"
)

type AppConfig struct {
	Conn       string
	Football   *usecases.FootballUsecase
	Automation *usecases.AutomationUseCase
	Process    *usecases.Process
	Logger     *log.CustomLog
}

type AppCache struct {
	PlayerModel *Players.PlayerRoot
}
