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
if len(database.DB.Matches)<=0 {
	fmt.Println("|     There's Still Nothing Here, Please Add Matches        |")
	fmt.Println("+-----------------------------------------------------------+")
}

	for i := 0; i < len(database.DB.Matches); i++ {
		match := database.DB.Matches[i]
		fmt.Printf("|%2d|%18s   %-3d  -  %3d   %-18s   | \n",match.MatchID, match.Team1.Name, match.Score1, match.Score2, match.Team2.Name)
		fmt.Println("+-----------------------------------------------------------+")
	}

}



func AddMatches() {
fmt.Println("=== TAMBAHKAN MATCH ===")
DisplayOnlyTeamsMenu()
	
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
	} else if team1Index==team2Index {
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

	database.DB.Matches = append(database.DB.Matches, match)

	fmt.Println("Match berhasil ditambahkan.")
	WinLoseDraw(match.Team1.Name, match.Team2.Name, score1, score2)
}

func EditMatches() {
	if len(database.DB.Matches)<=0{
		fmt.Println("Belum ada Match untuk diedit")
		return
	} else {
		ViewMatches()
		fmt.Println("=== Update Match Result ===")
		MatchID:= ScanNumber("Masukan ID Match yang ingin diubah: ")
		idx :=binarySearchByMatchID(MatchID)
		if idx == -1 {
			fmt.Println("Match Tidak Ditemukan")
			return
		} else {
			match:= database.DB.Matches[idx]
			fmt.Printf("Match Ditemukan \n")
			fmt.Printf("|%20s   %-3d  -  %3d   %-20s  | \n", match.Team1.Name, match.Score1, match.Score2, match.Team2.Name)
			database.DB.Matches[idx].Score1=ScanNumber("Masukan Skor Tim 1: ")
			database.DB.Matches[idx].Score2=ScanNumber("Masukan Skor Tim 2: ")
			match= database.DB.Matches[idx]
			fmt.Println("Match Berhasil di Update")
			fmt.Printf("|%20s   %-3d  -  %3d   %-20s  | \n", match.Team1.Name, match.Score1, match.Score2, match.Team2.Name)
			WinLoseDraw(match.Team1.Name, match.Team2.Name, match.Score1, match.Score2)
		
		}
	
		ScanString("Tekan Enter untuk melanjutkan...")
	}

}