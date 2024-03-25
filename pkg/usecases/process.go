package usecases

import (
	"DatabaseHandler/pkg/handlers/process"
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
	pr := m.Football.GetPlayerWithId(203)
	m.PlayerProcess.ProcessPlayer(pr)

}
