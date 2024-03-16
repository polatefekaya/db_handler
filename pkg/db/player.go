package db

import (
	"DatabaseHandler/pkg/data/models/Players"
	"DatabaseHandler/pkg/repository"
)

type PlayerDb struct {
	Repo repository.PlayerRepository
}

func (m *PlayerDb) AddPlayer(p *Players.PlayerRoot) {
	//calls repository.addplayer
	e := p.Responses[0].Player.ToEntity()
	m.Repo.AddPlayer(e)
	
}
