package models

type Database struct {
	Teams        []Team
	Classement   []Classement
	Players      []Player
	LastTeamID   int
	LastPlayerID int
}