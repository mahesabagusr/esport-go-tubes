package database

import (
	"esportgacor/models"
)

var DB *models.Database

func InitDB() *models.Database {
	teams, players, nTeam, nPlayer := GenerateDummyData()

	DB = &models.Database{
		Teams:        teams,
		NTeam:        nTeam,
		Classement:   [models.NMAX]models.Classement{},
		NClassement:  0,
		Players:      players,
		NPlayer:      nPlayer,
		Matches:      [models.NMAX]models.Match{},
		NMatch:       0,
		LastTeamID:   nTeam,
		LastPlayerID: nPlayer,
		LastMatchID:  0,
	}

	return DB
}