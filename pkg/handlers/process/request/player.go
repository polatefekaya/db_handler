package request

import (
	"DatabaseHandler/pkg/data/models/Players"
	"DatabaseHandler/pkg/usecases"
)

type PlayerProcessRequestHandler struct {
	FootballUseCase *usecases.FootballUsecase
}

func (m *PlayerProcessRequestHandler) GetPlayerWithId(id int) *Players.PlayerRoot {
	return m.FootballUseCase.GetPlayerWithId(id)
}
