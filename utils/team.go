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
	fmt.Println("Modify Team belum jadi rek")
}

func DeleteTeam(){
	fmt.Println("Delete Team Belum Jadi Rek")
}