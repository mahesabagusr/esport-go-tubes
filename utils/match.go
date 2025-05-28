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
	score1:= ScanNumber("Jumlah Score: ")
	team2:= ScanString("Nama Tim Kedua: ")
	score2:= ScanNumber("Jumlah Score: ")
	database.DB.LastMatchID++
	match := models.Match{
		MatchID: database.DB.LastMatchID,
		Team1:
		Score1: score1,
		Team2: team2, 
		Score2: score2,
	}
	database.DB.Matches= append(database.DB.Matches, match)
	
}

func EditMatches() {

}