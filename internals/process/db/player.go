package db

import (
	"DatabaseHandler/internals/process/data"
	e "DatabaseHandler/pkg/data/entities/Players"
)

type PlayerProcessDbHandler struct {
}

func (m *PlayerProcessDbHandler) StartDb(pe *e.PlayerEntity, se []*e.StatisticEntity, ts []*data.TempStatistic) {
	//make db calls
}
