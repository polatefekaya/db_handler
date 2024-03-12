package generics

import (
	"DatabaseHandler/pkg/data/entities"
	"DatabaseHandler/pkg/data/models"
	"reflect"
)

var EntityTypes = map[reflect.Type]reflect.Type{
	reflect.TypeOf(models.PlayerRoot{}): reflect.TypeOf(entities.PlayerEntity{}),
}

type EntityConvertTypes interface {
	models.PlayerRoot | models.Player
}
