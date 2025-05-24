package models

type Database struct {
	Teams        []Team
	Classement   []Classement
	Players      []Player
	Matches      []Match
	LastTeamID   int
	LastPlayerID int
	LastMatchID  int
}