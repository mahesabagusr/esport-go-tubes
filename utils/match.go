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
	
	idTeam1 := ScanNumber("ID Tim Pertama: ")
	team1Index := linearSearchByID(idTeam1)

	if team1Index == -1 {
		fmt.Println("Tim pertama tidak ditemukan.")
		return
	}

	score1 := ScanNumber("Jumlah Skor Tim Pertama: ")

	idTeam2 := ScanNumber("ID Tim Kedua: ")
	team2Index := linearSearchByID(idTeam2)
	if team2Index == -1 {
		fmt.Println("Tim kedua tidak ditemukan.")
		return
	}

	score2 := ScanNumber("Jumlah Skor Tim Kedua: ")

	database.DB.LastMatchID++
	match := models.Match{
		MatchID: database.DB.LastMatchID,
		Team1:   database.DB.Teams[team1Index],
		Score1:  score1,
		Team2:   database.DB.Teams[team2Index],
		Score2:  score2,
	}

	database.DB.Matches = append(database.DB.Matches, match)

	fmt.Println("Match berhasil ditambahkan.")
}

func EditMatches() {
	
}