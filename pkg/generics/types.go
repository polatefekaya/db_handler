package generics

import (
	"DatabaseHandler/pkg/data/entities/Players"
	Players2 "DatabaseHandler/pkg/data/models/Players"
	"reflect"
)

var EntityTypes = map[reflect.Type]reflect.Type{
	reflect.TypeOf(Players2.PlayerRoot{}): reflect.TypeOf(Players.PlayerEntity{}),
}

type EntityConvertTypes interface {
	Players2.PlayerRoot | Players2.Player
}
