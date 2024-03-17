package main

import (
	"DatabaseHandler/pkg/config"
	"DatabaseHandler/pkg/data/entities/Players"
	"DatabaseHandler/pkg/data/models/query"
	"DatabaseHandler/pkg/handlers"
	"DatabaseHandler/pkg/usecases"
	"log"
)

func main() {
	var app config.AppConfig

	app.Football = usecases.NewFootballUsecase()

	app.Automation = usecases.NewAutomationUseCase()
	app.Automation.FootballUsecase = app.Football

	req := app.Football.GetPlayerWithId(3)

	a := handlers.Generate(*req).(Players.PlayerEntity)

	b := req.Responses[0].Player.ToEntity()
	c := req.Responses[0].Statistics[0].Team.ToEntity()

	a.FirstName = req.Responses[0].Player.FirstName
	log.Println(b.Name)
	log.Println(c.Name)

	q := query.NewQuery()
	pl := q.NewPlayers()
	pl.Player.Generate("", "", "")

	qb := usecases.NewQueryBuilder()
	qb.Build(pl.Player)
}
