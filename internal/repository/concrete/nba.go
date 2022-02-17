package concrete

import (
	"fmt"
	"github.com/nbaSimulation/internal/models"
	"math/rand"
)

type MemoryDatabase struct {
	Statistics models.Statistic
	Players    []models.Player
}

func NewMemDB() *MemoryDatabase {
	return &MemoryDatabase{
		Statistics: models.Statistic{},
		Players:    []models.Player{},
	}
}

const attackCountInAMinute = 3
const homeTeamName = "TeamA"
const awayTeamName = "TeamB"

func (api *MemoryDatabase) CreateGame() {
	homeTeam := models.Team{
		Name: "TeamA",
	}
	awayTeam := models.Team{
		Name: "TeamB",
	}

	homeTeamPlayersLineup := []models.Player{
		{
			Name:         "PlayerA1",
			JerseyNumber: 1,
			Assists:      0,
			Points:       0,
			Team:         homeTeam,
		},
		{
			Name:         "PlayerA2",
			JerseyNumber: 2,
			Assists:      0,
			Points:       0,
			Team:         homeTeam,
		},
		{
			Name:         "PlayerA3",
			JerseyNumber: 3,
			Assists:      0,
			Points:       0,
			Team:         homeTeam,
		},
		{
			Name:         "PlayerA4",
			JerseyNumber: 4,
			Assists:      0,
			Points:       0,
			Team:         homeTeam,
		},
		{
			Name:         "PlayerA5",
			JerseyNumber: 5,
			Assists:      0,
			Points:       0,
			Team:         homeTeam,
		},
	}
	awayTeamPlayersLineup := []models.Player{
		{
			Name:         "PlayerB1",
			JerseyNumber: 1,
			Assists:      0,
			Points:       0,
			Team:         awayTeam,
		},
		{
			Name:         "PlayerB2",
			JerseyNumber: 2,
			Assists:      0,
			Points:       0,
			Team:         awayTeam,
		},
		{
			Name:         "PlayerB3",
			JerseyNumber: 3,
			Assists:      0,
			Points:       0,
			Team:         awayTeam,
		},
		{
			Name:         "PlayerB4",
			JerseyNumber: 4,
			Assists:      0,
			Points:       0,
			Team:         awayTeam,
		},
		{
			Name:         "PlayerB5",
			JerseyNumber: 5,
			Assists:      0,
			Points:       0,
			Team:         awayTeam,
		},
	}

	api.Statistics = models.Statistic{
		AwayTeam:      awayTeam,
		HomeTeam:      homeTeam,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		TopAssist:     models.Player{},
		TopScorer:     models.Player{},
	}

	api.Players = append(homeTeamPlayersLineup, awayTeamPlayersLineup...)

	fmt.Println(api.Statistics)
}

func (api *MemoryDatabase) UpdateScore() {
	possibleScores := []int{2, 3}

	for i := 0; i < attackCountInAMinute; i++ {
		player := api.Players[rand.Intn(len(api.Players))]
		assistPlayer := api.Players[rand.Intn(len(api.Players))]

		player.PointInAttack = possibleScores[rand.Intn(len(possibleScores))]
		assistPlayer.Assists++

		if player.Team.Name == homeTeamName {
			api.Statistics.HomeTeamScore += player.PointInAttack
		} else {
			api.Statistics.AwayTeamScore += player.PointInAttack
		}

		player.Points += player.PointInAttack

		for i, item := range api.Players {
			if item.Name == player.Name {
				api.Players[i] = player
			}

			if item.Name == assistPlayer.Name {
				api.Players[i] = assistPlayer
			}
		}

		api.Statistics.TopScorer = findTopScorer(api.Players)
		api.Statistics.TopAssist = findTopAssists(api.Players)
	}

	fmt.Printf("%s: %d, %s: %d, Top Score: %d, Top Scorer: %s, Top Assists: %d, Top Assist Player: %s",
		api.Statistics.HomeTeam,
		api.Statistics.HomeTeamScore,
		api.Statistics.AwayTeam,
		api.Statistics.AwayTeamScore,
		api.Statistics.TopScorer.Points,
		api.Statistics.TopScorer.Name,
		api.Statistics.TopAssist.Assists,
		api.Statistics.TopAssist.Name,
	)
	fmt.Println()
}

// region Helper Functions

func findTopScorer(players []models.Player) models.Player {
	topScorer := players[0]
	for _, player := range players {
		if player.Points > topScorer.Points {
			topScorer = player
		}
	}
	return topScorer
}

func findTopAssists(players []models.Player) models.Player {
	topAssists := players[0]
	for _, player := range players {
		if player.Assists > topAssists.Assists {
			topAssists = player
		}
	}
	return topAssists
}

// endregion
