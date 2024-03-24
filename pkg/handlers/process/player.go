package process

import (
	"DatabaseHandler/pkg/handlers/process/request"
)

type PlayerProcess struct {
	Db      IPlayerProcess
	Data    IPlayerProcess
	Request *request.PlayerProcessRequestHandler
}

type IPlayerProcess interface {
	PlayerProcess()
	TeamProcess()
	LeagueProcess()
}

func NewPlayerProcess() *PlayerProcess {
	return &PlayerProcess{}
}

func (m *PlayerProcess) ProcessPlayer(id int) {
	pr := m.Request.GetPlayerWithId(id)
	m.Data.StartData(pr)
	m.Data.LeagueProcess()
	m.Db.Start()
}
