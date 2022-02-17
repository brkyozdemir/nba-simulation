package models

type Player struct {
	Name          string
	JerseyNumber  int
	Assists       int
	Points        int
	Team          Team
	PointInAttack int
}
