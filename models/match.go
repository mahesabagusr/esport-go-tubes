package models

type Match struct {
	MatchID int
	Team1   Team
	Team2   Team
	Score1  int
	Score2  int
}