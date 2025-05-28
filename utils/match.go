package utils

import (
	"esportgacor/database"
	"esportgacor/models"
	"fmt"
)

func ViewMatches() {

fmt.Printf(`
+-----------------------------------------------------------+
|                      DAFTAR MATCH                         |
+-----------------------------------------------------------+
`)
	for i := 0; i < len(database.DB.Matches); i++ {
		match := database.DB.Matches[i]
		fmt.Printf("|%20s   %-3d  -  %3d   %-20s  | \n", match.Team1.Name, match.Score1, match.Score2, match.Team2.Name)
		fmt.Println("+-----------------------------------------------------------+")
	}
	ScanString("Tekan Enter untuk melanjutkan...")
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
	} else if team1Index==team2Index {
		fmt.Println("Tim tidak boleh sama")
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