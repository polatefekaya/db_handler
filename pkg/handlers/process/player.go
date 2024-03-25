package process

import (
	"DatabaseHandler/pkg/data/models/Players"
	log2 "DatabaseHandler/pkg/handlers/log"
	"DatabaseHandler/pkg/handlers/process/data"
	"DatabaseHandler/pkg/handlers/process/db"
	"log"
)

type PlayerProcess struct {
	Db   *db.PlayerProcessDbHandler
	Data *data.PlayerProcessDataHandler
}

var logger *log2.CustomLog

func NewPlayerProcess() *PlayerProcess {
	logger = log2.NewSLog()
	return &PlayerProcess{}
}

func (m *PlayerProcess) ProcessPlayer(root *Players.PlayerRoot) {
	pe, se, ts := m.Data.StartData(root)
	//m.Db.StartDb(pe, se, ts)
	logger.Log.Info("jshgkjsnhgkjnsfkjgnskjfngkjsfg")
	log.Println(pe.Id)
	log.Println(se[0].ShotId)
	log.Println(ts[0].TeamEntity.Name)
}
