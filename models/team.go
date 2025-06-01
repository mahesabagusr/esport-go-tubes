package models

type Team struct {
	ID      int
	Name    string
	Players [NMAX]Player
	NPlayer int 
	Coach   string
}

type Player struct {
	ID     int
	Name   string
	Kills  int
	Deaths int
}

type Classement struct {
	Team    Team
	Pts     int
	Win     int
	Lose    int
	Winrate int
}

