package config

import (
	"DatabaseHandler/pkg/data/models/Players"
	"DatabaseHandler/pkg/usecases"
)

type AppConfig struct {
	Conn       string
	Football   *usecases.FootballUsecase
	Automation *usecases.AutomationUseCase
}

type AppCache struct {
	PlayerModel *Players.PlayerRoot
}
