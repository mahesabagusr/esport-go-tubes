package models

type Team struct {
	id      int
	name    string
	players []Player
	coach   string
}

type Player struct {
	id     int
	name   string
	kills  int
	deaths int
}

type Classement struct {
	team []Team
	pts  int
	win int
	lose int
	winrate int
}