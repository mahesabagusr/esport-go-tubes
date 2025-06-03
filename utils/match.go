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
	if database.DB.NMatch <= 0 {
		fmt.Println("|     There's Still Nothing Here, Please Add Matches        |")
		fmt.Println("+-----------------------------------------------------------+")
		return
	}

	for i := 0; i < database.DB.NMatch; i++ {
		match := database.DB.Matches[i]
		fmt.Printf("|%2d|%18s   %-3d  -  %3d   %-18s   | \n", 
			match.MatchID, match.Team1.Name, match.Score1, match.Score2, match.Team2.Name)
		fmt.Println("+-----------------------------------------------------------+")
	}
}

func AddMatches() {
	fmt.Println("=== TAMBAHKAN MATCH ===")
	DisplayOnlyTeamsMenu()
	
	if database.DB.NMatch >= models.NMAX {
		fmt.Println("Tidak bisa menambahkan match lagi, kapasitas penuh!")
		return
	}

	idTeam1 := ScanNumber("ID Tim Pertama: ")
	team1Index := linearSearchByTeamID(idTeam1)
	if team1Index == -1 {
		fmt.Println("Tim pertama tidak ditemukan.")
		return
	}

	score1 := ScanNumber("Jumlah Skor Tim Pertama: ")

	idTeam2 := ScanNumber("ID Tim Kedua: ")
	team2Index := linearSearchByTeamID(idTeam2)
	if team2Index == -1 {
		fmt.Println("Tim kedua tidak ditemukan.")
		return
	} else if team1Index == team2Index {
		fmt.Println("Tim tidak boleh sama")
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


	database.DB.Matches[database.DB.NMatch] = match
	database.DB.NMatch++
  SortTeamsByPts(database.DB.Classement[:database.DB.NClassement])
	fmt.Println("Match berhasil ditambahkan.")
	WinLoseDraw(match.Team1.Name, match.Team2.Name, score1, score2, false)
}

func EditMatches() {
	if database.DB.NMatch <= 0 {
		fmt.Println("Belum ada Match untuk diedit")
		return
	}

	ViewMatches()
	fmt.Println("=== Update Match Result ===")
	MatchID := ScanNumber("Masukan ID Match yang ingin diubah: ")
	idx := linearSearchByTeamID(MatchID)
	if idx == -1 {
		fmt.Println("Match Tidak Ditemukan")
		return
	}

	match := database.DB.Matches[idx]
	fmt.Printf("Match Ditemukan \n")
	fmt.Printf("|%20s   %-3d  -  %3d   %-20s  | \n", 
	match.Team1.Name, match.Score1, match.Score2, match.Team2.Name)


	WinLoseDraw(match.Team1.Name, match.Team2.Name, match.Score1, match.Score2, true)


	newScore1 := ScanNumber("Masukan Skor Tim 1: ")
	newScore2 := ScanNumber("Masukan Skor Tim 2: ")


	database.DB.Matches[idx].Score1 = newScore1
	database.DB.Matches[idx].Score2 = newScore2


	WinLoseDraw(match.Team1.Name, match.Team2.Name, newScore1, newScore2, false)
	SortTeamsByPts(database.DB.Classement[:database.DB.NClassement])
	fmt.Println("Match Berhasil di Update")
	fmt.Printf("|%20s   %-3d  -  %3d   %-20s  | \n", 
		match.Team1.Name, newScore1, newScore2, match.Team2.Name)

	ScanString("Tekan Enter untuk melanjutkan...")
}