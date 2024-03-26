package usecases

import (
	"DatabaseHandler/internals/log"
	"DatabaseHandler/internals/process"
)

//make the request, process the data, upload to db

type Process struct {
	PlayerProcess *process.PlayerProcess
	Football      *FootballUsecase
}

func NewProcess() *Process {
	return &Process{
		PlayerProcess: process.NewPlayerProcess(),
	}
}

func (m *Process) Start() {
	log.INFO("Processing started")
	pr := m.Football.GetPlayerWithId(203)
	m.PlayerProcess.ProcessPlayer(pr)
	log.INFO("Processing ended")
}
