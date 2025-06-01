// utils/team_utils.go
package utils

import (
	"esportgacor/database"
	"esportgacor/models"
	"fmt"
	"strconv"
)

func AddTeam() {
	fmt.Println("=== Tambah Tim Baru ===")


	if database.DB.NTeam >= models.NMAX {
		fmt.Println("Tidak bisa menambahkan tim lagi, kapasitas penuh!")
		return
	}

	teamName := ScanString("Nama Tim: ")
	coachName := ScanString("Nama Coach: ")


	database.DB.LastTeamID++
	newTeam := models.Team{
		ID:      database.DB.LastTeamID,
		Name:    teamName,
		Coach:   coachName,
		NPlayer: 5, 
	}


	for i := 0; i < 5; i++ {
		fmt.Printf("\n-- Pemain #%d --\n", i+1)

		if database.DB.NPlayer >= models.NMAX {
			fmt.Println("Tidak bisa menambahkan pemain lagi, kapasitas penuh!")
			return
		}

		playerName := ScanString("Nama Pemain: ")
		kills := ScanNumber("Jumlah Kills: ")
		deaths := ScanNumber("Jumlah Deaths: ")

		database.DB.LastPlayerID++
		player := models.Player{
			ID:     database.DB.LastPlayerID,
			Name:   playerName,
			Kills:  kills,
			Deaths: deaths,
		}

		newTeam.Players[i] = player
		database.DB.Players[database.DB.NPlayer] = player
		database.DB.NPlayer++
	}


	database.DB.Teams[database.DB.NTeam] = newTeam
	database.DB.NTeam++

	fmt.Println("\nTim berhasil ditambahkan!")
}

func ModifyTeam() {
	fmt.Println("=== Update Tim ===")
	DisplayOnlyTeamsMenu()
	teamName := ScanString("Masukkan Nama Tim yang ingin diubah: ")


	idx := -1
	for i := 0; i < database.DB.NTeam; i++ {
		if database.DB.Teams[i].Name == teamName {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Tim tidak ditemukan!")
		return
	}

	DisplayOnlyTeamsAndPlayersMenu(idx)

	fmt.Println("=== Update Tim ===")
	newTeamName := ScanString("Nama Tim (Tekan Enter Jika Tidak Ingin Diubah): ")
	if newTeamName != "" {
		database.DB.Teams[idx].Name = newTeamName
	}

	coach := ScanString("Nama Coach (Tekan Enter Jika Tidak Ingin Diubah): ")
	if coach != "" {
		database.DB.Teams[idx].Coach = coach
	}


	for j := 0; j < database.DB.Teams[idx].NPlayer; j++ {
		fmt.Printf("\n-- Pemain #%d --\n", j+1)
		playerName := ScanString("Nama Pemain (Tekan Enter Jika Tidak Ingin Diubah): ")
		if playerName != "" {
			database.DB.Teams[idx].Players[j].Name = playerName
		}

		killsStr := ScanString("Jumlah Kills (Tekan Enter Jika Tidak Ingin Diubah): ")
		if killsStr != "" {
			kills, _ := strconv.Atoi(killsStr)
			database.DB.Teams[idx].Players[j].Kills = kills
		}

		deathsStr := ScanString("Jumlah Deaths (Tekan Enter Jika Tidak Ingin Diubah): ")
		if deathsStr != "" {
			deaths, _ := strconv.Atoi(deathsStr)
			database.DB.Teams[idx].Players[j].Deaths = deaths
		}
	}
}

func DeleteTeam() {
	var Team models.Team
	DisplayOnlyTeamsMenu()

	fmt.Println("=== Delete Tim ===")
	fmt.Println("||Match dan Ranking Klasemen Tidak Akan Ikut Terhapus||")
	teamName := ScanString("Masukkan Nama Tim yang ingin dihapus: ")


	idx := -1
	for i := 0; i < database.DB.NTeam; i++ {
		if database.DB.Teams[i].Name == teamName {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Tim tidak ditemukan!")
		return
	}


	for i := idx; i < database.DB.NTeam-1; i++ {
		database.DB.Teams[i] = database.DB.Teams[i+1]
	}


	database.DB.NTeam--

	database.DB.Teams[database.DB.NTeam] = Team

	fmt.Println("Tim berhasil dihapus!")
}