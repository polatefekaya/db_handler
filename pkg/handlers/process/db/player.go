package db

import (
	e "DatabaseHandler/pkg/data/entities/Players"
	"DatabaseHandler/pkg/handlers/process/data"
)

type PlayerProcessDbHandler struct {
}

func (m *PlayerProcessDbHandler) StartDb(pe *e.PlayerEntity, se []*e.StatisticEntity, ts []*data.TempStatistic) {
	//make db calls
}
