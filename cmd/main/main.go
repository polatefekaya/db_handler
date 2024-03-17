package main

import (
	"DatabaseHandler/pkg/config"
	"DatabaseHandler/pkg/usecases"
	env "github.com/joho/godotenv"
	"log"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	var app config.AppConfig
	Initialize(&app)

	app.Automation.AutomatePlayer()

}

func Initialize(app *config.AppConfig) {
	app.Football = usecases.NewFootballUsecase()

	app.Automation = usecases.NewAutomationUseCase()
	app.Automation.FootballUsecase = app.Football
	app.Automation.Process = usecases.NewProcess()
}
