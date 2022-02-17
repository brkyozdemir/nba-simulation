package cron

import (
	"fmt"
	"github.com/nbaSimulation/config"
	"github.com/nbaSimulation/internal/database"
	"strconv"
	"time"
)

type jobTicker struct {
	timer     *time.Timer
	totalTime time.Duration
}

const gameTime = 240

func BeginGame(api database.API) {
	jobTicker := &jobTicker{}
	fmt.Println("The game has began")
	cfg := config.GetConfig()
	seconds, _ := strconv.Atoi(cfg.Cron.TimeIntervalInSeconds)
	jobTicker.updateTimer(seconds)
	api.Db.CreateGame()
	for {
		<-jobTicker.timer.C
		api.Db.UpdateScore()
		jobTicker.updateTimer(seconds)
		jobTicker.totalTime += time.Second * time.Duration(seconds) * 12
		if jobTicker.totalTime == time.Second*time.Duration(gameTime) {
			jobTicker.timer.Stop()
		}

		fmt.Println("Game Time: ", jobTicker.totalTime)
	}
}

func (t *jobTicker) updateTimer(seconds int) {
	diff := time.Second * time.Duration(seconds)

	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}
