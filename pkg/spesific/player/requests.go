package player

import (
	"DatabaseHandler/pkg/handlers"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetPlayerWithId(id int) *PlayerRoot {
	body := handlers.GetRequest("d")

	var pr PlayerRoot
	handlers.FromJson(body, &pr)

	return &pr
}

func GetPlayerWithName(name string) *PlayerRoot {
	pr := PlayerRoot{
		responses: nil,
	}
	return &pr
}

func GetPlayersWithTeamId(id int) *PlayerRoot {
	pr := PlayerRoot{
		responses: nil,
	}
	return &pr
}

func dummy() {
	//create new request
	url := "http:ss"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//add headers
	req.Header.Add("X-Header", "val")

	//make the call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//convert it to json
	var um PlayerRoot
	err = json.Unmarshal(body, &um)
	if err != nil {
		log.Fatal(err)
	}
}
