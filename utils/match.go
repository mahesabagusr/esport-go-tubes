package utils

import (
	"esportgacor/database"
	"esportgacor/models"
	"fmt"
)

func ViewMatches() {

}

func AddMatches() {
	fmt.Println("=== TAMBAHKAN MATCH ===")
	team1:= ScanString("Nama Tim Pertama: ")

	idx := -1
	for i := 0; i < len(database.DB.Teams); i++ {
		if database.DB.Teams[i].Name == team1 {
			idx = i
		}
	}

	score1:= ScanNumber("Jumlah Score: ")

	database.DB.LastMatchID++
	match := models.Match{
		MatchID: database.DB.LastMatchID,
		Team1: database.DB.Teams[idx], 
		Score1: score1,
		Team2: team2, 
		Score2: score2,
	}
	database.DB.Matches= append(database.DB.Matches, match)
	
}

func EditMatches() {

}