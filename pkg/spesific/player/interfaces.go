package player

import (
	"DatabaseHandler/pkg/data/entities/Players"
	Players2 "DatabaseHandler/pkg/data/models/Players"
)

type IPlayer interface {
	ToPlayerEntity(pr *Players2.PlayerRoot) *Players.PlayerEntity
}
