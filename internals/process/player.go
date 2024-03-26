package process

import (
	log2 "DatabaseHandler/internals/log"
	"DatabaseHandler/internals/process/data"
	"DatabaseHandler/internals/process/db"
	"DatabaseHandler/pkg/data/models/Players"
	"log"
)

type PlayerProcess struct {
	Db   *db.PlayerProcessDbHandler
	Data *data.PlayerProcessDataHandler
}

func NewPlayerProcess() *PlayerProcess {
	return &PlayerProcess{}
}

func (m *PlayerProcess) ProcessPlayer(root *Players.PlayerRoot) {
	log2.INFO("General player process started")
	pe, se, ts := m.Data.StartData(root)
	log2.INFO("General player process ended with, ", "Player", pe.Id, "Statistics", se)
	//m.Db.StartDb(pe, se, ts)
	log.Println(ts[0].TeamEntity.Name)
}
