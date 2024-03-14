package main

import (
	"DatabaseHandler/pkg/config"
	"DatabaseHandler/pkg/data/entities"
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

	a := handlers.Generate(*req).(entities.PlayerEntity)
	a.FirstName = req.Responses[0].Player.FirstName
	log.Println(a.FirstName)

}
