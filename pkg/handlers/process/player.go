package process

import (
	"DatabaseHandler/pkg/handlers/process/data"
	"DatabaseHandler/pkg/handlers/process/db"
	"DatabaseHandler/pkg/handlers/process/request"
)

type PlayerProcess struct {
	Db      *db.PlayerProcessDbHandler
	Data    *data.PlayerProcessDataHandler
	Request *request.PlayerProcessRequestHandler
}

func NewPlayerProcess() *PlayerProcess {
	return &PlayerProcess{}
}

func (m *PlayerProcess) ProcessPlayer(id int) {
	pr := m.Request.GetPlayerWithId(id)
	m.Data.StartData(pr)
	m.Db.StartDb()
}
