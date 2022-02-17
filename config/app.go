package config

import "github.com/nbaSimulation/config/helpers"

type App struct {
	Name    string
	Env     string
	Port    string
	BaseUrl string
}

func NewApp() App {
	return App{
		Name:    helpers.Getenv("APP_NAME", "nba-simulation"),
		Env:     helpers.Getenv("APP_ENV", "local"),
		Port:    helpers.Getenv("PORT", "9000"),
		BaseUrl: helpers.Getenv("BASE_URL", "http://localhost"),
	}
}
