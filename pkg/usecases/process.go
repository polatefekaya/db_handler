package usecases

import (
	"DatabaseHandler/pkg/handlers/process"
)

//make the request, process the data, upload to db

type Process struct {
	PlayerProcess   *process.PlayerProcess
	FootballUsecase *FootballUsecase
}

func NewProcess() *Process {
	return &Process{
		PlayerProcess:   process.NewPlayerProcess(),
		FootballUsecase: NewFootballUsecase(),
	}
}

func (m *Process) Start() {
	pr := m.FootballUsecase.GetPlayerWithId(203)
	m.PlayerProcess.ProcessPlayer(pr)

}
