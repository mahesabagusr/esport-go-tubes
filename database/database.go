package database

import (
	"esportgacor/models"
)

var DB *models.Database

func InitDB() *models.Database {
	teams, classement, players:= GenerateDummyData()

	lastTeamID := len(teams)
	lastPlayerID := len(players)

	DB = &models.Database{
		Teams:        teams,
		Classement:   classement,
		Players:      players,
		Matches:      []models.Match{},
		LastTeamID:   lastTeamID,
		LastPlayerID: lastPlayerID,

	}

	return DB
}
