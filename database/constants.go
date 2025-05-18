package database

import "esportgacor/models"


func GenerateDummyData() ([]models.Team, []models.Classement, []models.Player) {
	players := []models.Player{
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

	teams := []models.Team{
		{
			ID:      1,
			Name:    "T1",
			Players: []models.Player{players[0], players[3], players[10], players[8], players[14]},
			Coach:   "Bengi",
		},
		{
			ID:      2,
			Name:    "Fnatic",
			Players: []models.Player{players[4], players[5], players[1], players[6], players[7]},
			Coach:   "YamatoCannon",
		},
		{
			ID:      3,
			Name:    "DWG KIA",
			Players: []models.Player{players[9], players[10], players[11], players[12], players[13]},
			Coach:   "kkOma",
		},
		{
			ID:      4,
			Name:    "G2 Esports",
			Players: []models.Player{players[2], players[4], players[5], players[8], players[14]},
			Coach:   "GrabbZ",
		},
		{
			ID:      5,
			Name:    "Team Liquid",
			Players: []models.Player{players[6], players[7], players[3], players[0], players[5]},
			Coach:   "Jatt",
		},
	}

	classement := []models.Classement{
		{
			Team:    []models.Team{teams[0]},
			Pts:     45,
			Win:     15,
			Lose:    2,
			Winrate: 88,
		},
		{
			Team:    []models.Team{teams[2]},
			Pts:     42,
			Win:     14,
			Lose:    3,
			Winrate: 82,
		},
		{
			Team:    []models.Team{teams[3]},
			Pts:     38,
			Win:     12,
			Lose:    5,
			Winrate: 71,
		},
		{
			Team:    []models.Team{teams[1]},
			Pts:     35,
			Win:     11,
			Lose:    6,
			Winrate: 65,
		},
		{
			Team:    []models.Team{teams[4]},
			Pts:     30,
			Win:     10,
			Lose:    7,
			Winrate: 59,
		},
	}

	return teams, classement, players
}