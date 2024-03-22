package usecases

import (
	"DatabaseHandler/pkg/handlers/process"
)

//make the request, process the data, upload to db

type Process struct {
	PlayerProcess *process.PlayerProcess
}

func NewProcess() *Process {
	return &Process{
		PlayerProcess: process.NewPlayerProcess(),
	}
}

func (m *Process) Start() {

	m.PlayerProcess.ProcessPlayer(404)

}
