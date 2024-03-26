package main

import (
	clog "DatabaseHandler/internals/log"
	"DatabaseHandler/pkg/config"
	"DatabaseHandler/pkg/usecases"
	env "github.com/joho/godotenv"
	"log"
)

func main() {
	clog.INFO("Application started")

	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}
	clog.INFO("Environment loaded")

	var app config.AppConfig
	Initialize(&app)

	app.Automation.AutomatePlayer()
}

func Initialize(app *config.AppConfig) {
	app.Football = usecases.NewFootballUsecase()

	app.Automation = usecases.NewAutomationUseCase()
	app.Automation.Process = usecases.NewProcess()
	app.Automation.Process.Football = app.Football

	app.Process = usecases.NewProcess()

	app.Logger = clog.NewSLog()
}
