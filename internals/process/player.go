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
	pe, se, ts := m.Data.StartData(root)
	//m.Db.StartDb(pe, se, ts)
	log2.WithUsers().INFO("DENEME")
	log.Println(pe.Id)
	log.Println(se[0].ShotId)
	log.Println(ts[0].TeamEntity.Name)
}
