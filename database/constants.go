package database

import "esportgacor/models"

func GenerateDummyData() ([models.NMAX]models.Team, [models.NMAX]models.Player, int, int) {
	var players [models.NMAX]models.Player
	var teams [models.NMAX]models.Team


	playerData := []models.Player{
		{ID: 1, Name: "Faker", Kills: 287, Deaths: 56},
		{ID: 2, Name: "Uzi", Kills: 245, Deaths: 72},
		{ID: 3, Name: "Rookie", Kills: 198, Deaths: 64},
		{ID: 4, Name: "TheShy", Kills: 176, Deaths: 83},
		{ID: 5, Name: "Caps", Kills: 203, Deaths: 71},
		{ID: 6, Name: "Perkz", Kills: 187, Deaths: 68},
		{ID: 7, Name: "Doublelift", Kills: 165, Deaths: 92},
		{ID: 8, Name: "CoreJJ", Kills: 42, Deaths: 107},
		{ID: 9, Name: "JackeyLove", Kills: 231, Deaths: 65},
		{ID: 10, Name: "Nuguri", Kills: 156, Deaths: 78},
		{ID: 11, Name: "ShowMaker", Kills: 212, Deaths: 59},
		{ID: 12, Name: "BeryL", Kills: 37, Deaths: 113},
		{ID: 13, Name: "Chovy", Kills: 225, Deaths: 63},
		{ID: 14, Name: "Deft", Kills: 189, Deaths: 74},
		{ID: 15, Name: "Keria", Kills: 51, Deaths: 98},
	}


	nPlayer := len(playerData)
	for i := 0; i < nPlayer; i++ {
		players[i] = playerData[i]
	}


	teamData := []models.Team{
		{
			ID:      1,
			Name:    "T1",
			Coach:   "Bengi",
		},
		{
			ID:      2,
			Name:    "Fnatic",
			Coach:   "YamatoCannon",
		},
		{
			ID:      3,
			Name:    "DWG KIA",
			Coach:   "kkOma",
		},
		{
			ID:      4,
			Name:    "G2 Esports",
			Coach:   "GrabbZ",
		},
		{
			ID:      5,
			Name:    "Team Liquid",
			Coach:   "Jatt",
		},
	}


	nTeam := len(teamData)
	for i := 0; i < nTeam; i++ {
		teams[i] = teamData[i]
		

		switch i {
		case 0: 
			teams[i].Players[0] = players[0]
			teams[i].Players[1] = players[3]
			teams[i].Players[2] = players[10]
			teams[i].Players[3] = players[8]
			teams[i].Players[4] = players[14]
			teams[i].NPlayer = 5
		case 1: 
			teams[i].Players[0] = players[4]
			teams[i].Players[1] = players[5]
			teams[i].Players[2] = players[1]
			teams[i].Players[3] = players[6]
			teams[i].Players[4] = players[7]
			teams[i].NPlayer = 5
		case 2: 
			teams[i].Players[0] = players[9]
			teams[i].Players[1] = players[10]
			teams[i].Players[2] = players[11]
			teams[i].Players[3] = players[12]
			teams[i].Players[4] = players[13]
			teams[i].NPlayer = 5
		case 3: 
			teams[i].Players[0] = players[2]
			teams[i].Players[1] = players[4]
			teams[i].Players[2] = players[5]
			teams[i].Players[3] = players[8]
			teams[i].Players[4] = players[14]
			teams[i].NPlayer = 5
		case 4: 
			teams[i].Players[0] = players[6]
			teams[i].Players[1] = players[7]
			teams[i].Players[2] = players[3]
			teams[i].Players[3] = players[0]
			teams[i].Players[4] = players[5]
			teams[i].NPlayer = 5
		}
	}

	return teams, players, nTeam, nPlayer
}