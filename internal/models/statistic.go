package models

type Statistic struct {
	HomeTeamScore int
	AwayTeamScore int
	HomeTeam      Team
	AwayTeam      Team
	TopAssist     Player
	TopScorer     Player
}
