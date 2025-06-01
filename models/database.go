package models

const NMAX = 1000


type Database struct {
	Teams        [NMAX]Team
	NTeam        int 
	Classement   [NMAX]Classement
	NClassement  int 
	Players      [NMAX]Player
	NPlayer      int 
	Matches      [NMAX]Match
	NMatch       int 
	LastTeamID   int
	LastPlayerID int
	LastMatchID  int
}