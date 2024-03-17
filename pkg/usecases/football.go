package usecases

import (
	m "DatabaseHandler/pkg/data/models/Players"
	q "DatabaseHandler/pkg/data/models/query"
	h "DatabaseHandler/pkg/handlers"
	"log"
	"os"
	sc "strconv"
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
	query, err := pq.Generate(sc.Itoa(id), sc.Itoa(2023), "").Build()
	if err != nil {
		log.Fatal(err)
	}

	key := os.Getenv("SPORTS_API_KEY")

	sa := h.CreateSportsApi(key, query)
	log.Println("Query", query)

	body := h.GetRequest(sa)

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
