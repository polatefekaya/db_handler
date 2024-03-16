package usecases

import (
	m "DatabaseHandler/pkg/data/models/Players"
	h "DatabaseHandler/pkg/handlers"
	"fmt"
)

type FootballUsecase struct{}

func NewFootballUsecase() *FootballUsecase {
	return &FootballUsecase{}
}

func (f *FootballUsecase) GetPlayerWithId(id int) *m.PlayerRoot {
	sa := h.CreateSportsApi("faa42408eae63bf0cf0dfb0ff4e1678d",
		fmt.Sprintf("players?id=%d&season=%d", id, 2023))

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
