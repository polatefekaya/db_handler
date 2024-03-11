package player

import (
	"DatabaseHandler/pkg/handlers"
)

func GetPlayerWithId(id int) *PlayerRoot {
	sa := handlers.CreateSportsApi("faa42408eae63bf0cf0dfb0ff4e1678d", "players?id=276&season=2019")

	body := handlers.GetRequest(sa)

	//faa42408eae63bf0cf0dfb0ff4e1678d
	var pr PlayerRoot
	handlers.FromJson(body, &pr)

	return &pr
}

func GetPlayerWithName(name string) *PlayerRoot {
	pr := PlayerRoot{
		//responses: nil,
	}
	return &pr
}

func GetPlayersWithTeamId(id int) *PlayerRoot {
	pr := PlayerRoot{
		//responses: nil,
	}
	return &pr
}
