package database

import (
	"esportgacor/models"
)

var DB *models.Database

func InitDB() *models.Database {
	teams, players:= GenerateDummyData()

	lastTeamID := len(teams)
	lastPlayerID := len(players)
	lastMatchID := 0

	DB = &models.Database{
		Teams:         teams,
		Classement:    []models.Classement{},
		Players:       players,
		Matches: 	   []models.Match{},
		LastTeamID:    lastTeamID,
		LastPlayerID:  lastPlayerID,
		LastMatchID:   lastMatchID,
	}

	return DB
}
