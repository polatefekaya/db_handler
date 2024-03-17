package usecases

import (
	m "DatabaseHandler/pkg/data/models/Players"
	q "DatabaseHandler/pkg/data/models/query"
	h "DatabaseHandler/pkg/handlers"
	"log"
	"os"
	"strconv"
)

type FootballUsecase struct {
	Query *q.Player
}

var pq *q.Player

func NewFootballUsecase() *FootballUsecase {
	pq = q.NewPlayer()
	return &FootballUsecase{}
}

func (f *FootballUsecase) GetPlayerWithId(id int) *m.PlayerRoot {
	query, err := pq.Generate(strconv.Itoa(id), strconv.Itoa(2023), "").Build()
	if err != nil {
		log.Fatal(err)
	}
	key := os.Getenv("SportsApiKey")
	log.Println(key)
	sa := h.CreateSportsApi(key, query)
	log.Println("Query", query)
	body := h.GetRequest(sa)

	//faa42408eae63bf0cf0dfb0ff4e1678d
	var pr m.PlayerRoot
	h.FromJson(body, &pr)

	return &pr
}

func GetPlayerWithName(name string) *m.PlayerRoot {
	pr := m.PlayerRoot{
		//responses: nil,
	}
	return &pr
}

func GetPlayersWithTeamId(id int) *m.PlayerRoot {
	pr := m.PlayerRoot{
		//responses: nil,
	}
	return &pr
}
