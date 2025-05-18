package utils

import (
	"esportgacor/database"
	"esportgacor/models"
	"fmt"
)

func AddTeam() {
	fmt.Println("=== Tambah Tim Baru ===")

	var teamName, coachName string
	fmt.Print("Nama Tim: ")
	fmt.Scan(&teamName)

	fmt.Print("Nama Coach: ")
	fmt.Scan(&coachName)

	var players []models.Player
	for i := 0; i < 5; i++ {
		fmt.Printf("\n-- Pemain #%d --\n", i+1)

		var playerName string
		var kills, deaths int

		fmt.Print("Nama Pemain: ")
		fmt.Scan(&playerName)

		fmt.Print("Jumlah Kills: ")
		fmt.Scan(&kills)

		fmt.Print("Jumlah Deaths: ")
		fmt.Scan(&deaths)

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

	fmt.Println("\n✅ Tim berhasil ditambahkan!")
}


func Modifyteam(){
	fmt.Println("=== Update Tim ===")
	var teamName, coach, player,kills,death string = "", "", "", "", ""
	fmt.Print("Masukkan Nama Tim yang ingin diubah: ")
	fmt.Scan(&teamName)
	
	for i := 0 ; i < len(database.DB.Teams); i++ {
		if database.DB.Teams[i].Name != teamName {
			fmt.Println("Tim tidak ditemukan")
			return
		}

			fmt.Println("=== Update Tim ===")
			fmt.Print("Nama Tim (Isi '-' Jika Tidak Ingin Diubah): ")
			fmt.Scan(&teamName)

			if coach != "-" {
				database.DB.Teams[i].Name = teamName
			}

			fmt.Print("Nama Coach (Isi '-' Jika Tidak Ingin Diubah): ")
			fmt.Scan(&database.DB.Teams[i].Coach)

			if coach != "-" {
				database.DB.Teams[i].Coach = coach
			}

			for j := 0; j < len(database.DB.Teams[i].Players); j++ {
				fmt.Printf("\n-- Pemain #%d --\n", j+1)
				fmt.Print("Nama Pemain (Isi '-' Jika Tidak Ingin Diubah): ")
				fmt.Scan(&player)

				if player != "-" {
					database.DB.Teams[i].Players[j].Name = player
				}

				fmt.Print("Jumlah Kills (Isi '-' Jika Tidak Ingin Diubah): ")
				fmt.Scan(&kills)

				if kills != "-" {
					database.DB.Teams[i].Players[j].Name = kills
				}

				fmt.Print("Jumlah Deaths (Isi '-' Jika Tidak Ingin Diubah): ")
				fmt.Scan(&death)

				if death != "-" {
					database.DB.Teams[i].Players[j].Name = death
				}
			}
	
	}

}

func DeleteTeam(){
	fmt.Println("Delete Team Belum Jadi Rek")
}