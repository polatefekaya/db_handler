package main

import (
	"DatabaseHandler/pkg/config"
	"DatabaseHandler/pkg/usecases"
	"log"
)

func main() {
	var app config.AppConfig

	app.Football = usecases.NewFootballUsecase()

	app.Automation = usecases.NewAutomationUseCase()
	app.Automation.FootballUsecase = app.Football

	req := app.Football.GetPlayerWithId(30)

	b := req.Responses[0].Player.ToEntity()
	c := req.Responses[0].Statistics[0].Team.ToEntity()

	log.Println(b.Name)
	log.Println(c.Name)

}
