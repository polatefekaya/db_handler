package process

import (
	"DatabaseHandler/pkg/data/models/Players"
	"DatabaseHandler/pkg/handlers/process/data"
	"DatabaseHandler/pkg/handlers/process/db"
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
	log.Println(pe.Id)
	log.Println(se[0].ShotId)
	log.Println(ts[0].TeamEntity.Name)
}
