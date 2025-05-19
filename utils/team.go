package utils

import (
	"esportgacor/database"
	"esportgacor/models"
	"fmt"
	"strconv"
)

func AddTeam() {
	fmt.Println("=== Tambah Tim Baru ===")

	teamName := ScanString("Nama Tim: ")

	coachName := ScanString("Nama Coach: ")

	var players []models.Player
	for i := 0; i < 5; i++ {
		fmt.Printf("\n-- Pemain #%d --\n", i+1)

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
		players = append(players, player)
		database.DB.Players = append(database.DB.Players, player)
	}

	database.DB.LastTeamID++
	newTeam := models.Team{
		ID:      database.DB.LastTeamID,
		Name:    teamName,
		Coach:   coachName,
		Players: players,
	}

	database.DB.Teams = append(database.DB.Teams, newTeam)

	fmt.Println("\nTim berhasil ditambahkan!")
}

func Modifyteam() {
	fmt.Println("=== Update Tim ===")
	DisplayOnlyTeamsMenu()
	teamName := ScanString("Masukkan Nama Tim yang ingin diubah: ")

	for i := 0; i < len(database.DB.Teams); i++ {
		if database.DB.Teams[i].Name != teamName {
			fmt.Println("Tim tidak ditemukan")
			return
		}

		DisplayOnlyTeamsAndPlayersMenu(i)

		fmt.Println("=== Update Tim ===")
		newTeamName := ScanString("Nama Tim (Tekan Enter Jika Tidak Ingin Diubah): ")

		if newTeamName != "" {
			database.DB.Teams[i].Name = newTeamName
		}

		coach := ScanString("Nama Coach (Tekan Enter Jika Tidak Ingin Diubah): ")

		if coach != "" {
			database.DB.Teams[i].Coach = coach
		}

		for j := 0; j < len(database.DB.Teams[i].Players); j++ {
			fmt.Printf("\n-- Pemain #%d --\n", j+1)
			player := ScanString("Nama Pemain (Tekan Enter Jika Tidak Ingin Diubah): ")

			if player != "" {
				database.DB.Teams[i].Players[j].Name = player
			}

			killsStr := ScanString("Jumlah Kills (Tekan Enter Jika Tidak Ingin Diubah): ")
			if killsStr != "" {
				kills, _ := strconv.Atoi(killsStr)
				database.DB.Teams[i].Players[j].Kills = kills
			}

			deathsStr := ScanString("Jumlah Deaths (Tekan Enter Jika Tidak Ingin Diubah): ")
			if deathsStr != "" {
				deaths, _ := strconv.Atoi(deathsStr)
				database.DB.Teams[i].Players[j].Deaths = deaths
			}
		}
	}
}

func DeleteTeam() {
	sortedTeam := SortTeamsWithName(database.DB.Teams)

	fmt.Println("=== Delete Tim ===")
	teamName := ScanString("Masukkan Nama Tim yang ingin dihapus: ")

	found := false
	var targetTeamId int = -1

	low, high := 0, len(sortedTeam)-1

	for low <= high {
		mid := low + (high-low)/2
		if sortedTeam[mid].Name == teamName {
			found = true
			targetTeamId = sortedTeam[mid].ID
			break
		} else if sortedTeam[mid].Name < teamName {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		i := 0
		indexToDelete := -1
		for i < len(database.DB.Teams) {
			if database.DB.Teams[i].ID == targetTeamId {
				indexToDelete = i
				break
			}
			i++
		}
		database.DB.Teams = append(database.DB.Teams[:indexToDelete], database.DB.Teams[indexToDelete+1:]...)
		fmt.Println("Tim berhasil dihapus!")
	} else {
		fmt.Println("Tim tidak ditemukan!")
	}
}