package database

import (
	"esportgacor/models"
)

var DB *models.Database

func InitDB() *models.Database {
	teams, classement:= GenerateDummyData()

	var allPlayers []models.Player
	for _, team := range teams {
		allPlayers = append(allPlayers, team.Players...)
	}

	lastTeamID := 0
	lastPlayerID := 0
	for _, team := range teams {
		if team.ID > lastTeamID {
			lastTeamID = team.ID
		}
		for _, player := range team.Players {
			if player.ID > lastPlayerID {
				lastPlayerID = player.ID
			}
		}
	}

	DB = &models.Database{
		Teams:        teams,
		Classement:   classement,
		Players:      allPlayers,
		LastTeamID:   lastTeamID,
		LastPlayerID: lastPlayerID,
	}

	return DB
}
