package main

import (
	"github.com/nbaSimulation/internal/cron"
	"github.com/nbaSimulation/internal/database"
	"github.com/nbaSimulation/internal/repository/concrete"
)

func main() {
	db := concrete.NewMemDB()

	api := database.API{
		Db: db,
	}

	cron.BeginGame(api)
}
