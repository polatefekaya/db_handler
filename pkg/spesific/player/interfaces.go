package player

import (
	"DatabaseHandler/pkg/data/entities"
	"DatabaseHandler/pkg/data/models"
)

type IPlayer interface {
	ToPlayerEntity(pr *models.PlayerRoot) *entities.PlayerEntity
}
