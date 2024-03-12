package handlers

import (
	"DatabaseHandler/pkg/data/entities"
	"DatabaseHandler/pkg/data/models"
	"DatabaseHandler/pkg/generics"
	"log"
	"reflect"
)

// Get model give entity
func Sample() {
	a := Generate(models.PlayerRoot{}).(entities.PlayerEntity)
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
