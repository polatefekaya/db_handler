package handlers

import (
	"encoding/json"
	"log"
)

func FromJson(b *[]byte, v any) {
	//convert it to json
	err := json.Unmarshal(*b, v)
	if err != nil {
		log.Fatal(err)
	}
}
