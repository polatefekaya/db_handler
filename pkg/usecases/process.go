package usecases

import (
	e "DatabaseHandler/pkg/data/entities/Players"
	"DatabaseHandler/pkg/data/models/Players"
	"log"
)

//make the request, process the data, upload to db

type Process struct {
	FootballUseCase *FootballUsecase
}

func NewProcess() *Process {
	return &Process{}
}

func (m *Process) Start() {

	m.startGetPlayerWithId(203)

}

func (m *Process) startGetPlayerWithId(id int) {
	//make request
	pr := m.FootballUseCase.GetPlayerWithId(id)

	//process data
	pe := pr.Responses[0].Player.ToEntity()
	log.Println(pe.Name)
	log.Println(pe.Id)
	//upload to db
	m.startDb(pe)
}

func (m *Process) startPlayerRequest() *Players.PlayerRoot {
	return m.FootballUseCase.GetPlayerWithId(200)
}

func (m *Process) startPlayerData(root *Players.PlayerRoot) *e.PlayerEntity {
	return root.Responses[0].Player.ToEntity()
}

func (m *Process) startDb(entity *e.PlayerEntity) {

}
