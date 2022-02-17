package config

import "github.com/nbaSimulation/config/helpers"

type Cron struct {
	TimeIntervalInSeconds string
}

func NewCron() Cron {
	return Cron{
		TimeIntervalInSeconds: helpers.Getenv("TIME_INTERVAL_IN_MINUTES", "5"),
	}
}
