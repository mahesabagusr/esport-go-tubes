package database

import (
	"esportgacor/models"
)

var DB *models.Database

func InitDB() *models.Database {
	teams, classement, players:= GenerateDummyData()

	lastTeamID := len(teams)
	lastPlayerID := len(players)
	lastMatchID := len(DB.Matches)

	DB = &models.Database{
		Teams:        teams,
		Classement:   classement,
		Players:      players,
		Matches: 		[]models.Match{},
		LastTeamID:   lastTeamID,
		LastPlayerID: lastPlayerID,
		LastMatchID: lastMatchID,
	}

	return DB
}
