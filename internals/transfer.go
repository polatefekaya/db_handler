package internals

import (
	"DatabaseHandler/internals/generics"
	"DatabaseHandler/pkg/data/entities/Players"
	Players2 "DatabaseHandler/pkg/data/models/Players"
	"log"
	"reflect"
)

// Get model give entity
func Sample() {
	a := Generate(Players2.PlayerRoot{}).(Players.PlayerEntity)
	log.Println(a.FirstName)

}

func Generate[T generics.EntityConvertTypes](i T) interface{} {
	a := generics.EntityTypes[reflect.TypeOf(i)]
	if a == nil {
		log.Fatal("Given parameter can not be a nil value")
	}
	b := reflect.New(a).Elem()
	if b.Kind() != reflect.Struct {
		log.Fatal("Kinds not matching")
	}
	return b.Interface()
}
