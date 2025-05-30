package utils

import (
	"esportgacor/database"
	"fmt"
)

func DisplayMainMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                       E SPORT GACOR                       |
+-----------------------------------------------------------+
| 1. MANAGE TEAM                                            |
| 2. DISPLAY ALL TEAMS                                      |
| 3. DISPLAY CLASSEMENT                                     |
| 4. MANAGE MATCHES											|
+-----------------------------------------------------------+
`)
}

func DisplayManageTeamMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                       MANAGE TEAM                         |
+-----------------------------------------------------------+
| 1. ADD TEAM                                               |
| 2. MODIFY TEAM                                            |
| 3. DELETE TEAM                                            |
| 4. BACK                                            	    |
+-----------------------------------------------------------+
`)
}

func DisplayManageMatchMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                      MANAGE MATCH                         |
+-----------------------------------------------------------+
| 1. ADD MATCH                                              |
| 2. MODIFY MATCH                                           |
| 3. VIEW MATCH                                             |
| 4. BACK                                            	    |
+-----------------------------------------------------------+
`)
}

func DisplayAllTeamsAndPlayersMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                   DAFTAR SEMUA TIM & PEMAIN               |
+-----------------------------------------------------------+
`)
	for i := 0; i < len(database.DB.Teams); i++ {
		team := database.DB.Teams[i]
		fmt.Printf("| ID: %-3d Nama Tim: %-20s Coach: %-10s |\n", team.ID, team.Name, team.Coach)
		fmt.Println("| Pemain:")
		for j := 0; j < len(team.Players); j++ {
			player := team.Players[j]
			fmt.Printf("|   - ID: %-3d Nama: %-15s Kills: %-3d Deaths: %-3d |\n", player.ID, player.Name, player.Kills, player.Deaths)
		}
		fmt.Println("+-----------------------------------------------------------+")
	}
	ScanString("Tekan Enter untuk melanjutkan...")
}

func DisplayOnlyTeamsMenu(){
	    fmt.Printf(`
+---------------------------+
|       DAFTAR TIM          |
+---------------------------+
`)
    for i := 0; i < len(database.DB.Teams); i++ {
        team := database.DB.Teams[i]
        fmt.Printf("| %-3d | %-20s |\n", team.ID, team.Name)
    }
    fmt.Println("+---------------------------+")
}

func DisplayOnlyTeamsAndPlayersMenu(i int) {
	fmt.Printf(`
+-----------------------------------------------------------+
|                   DAFTAR SEMUA TIM & PEMAIN               |
+-----------------------------------------------------------+
`)
		team := database.DB.Teams[i]
		fmt.Printf("| ID: %-3d Nama Tim: %-20s Coach: %-10s |\n", team.ID, team.Name, team.Coach)
		fmt.Println("| Pemain:")
		for j := 0; j < len(team.Players); j++ {
			player := team.Players[j]
			fmt.Printf("|   - ID: %-3d Nama: %-15s Kills: %-3d Deaths: %-3d |\n", player.ID, player.Name, player.Kills, player.Deaths)
		}
		fmt.Println("+-----------------------------------------------------------+")
}

func DisplayClassement(){
	fmt.Printf(`
+-----------------------------------------------------------+
|                         KLASEMEN                          |
+-----------------------------------------------------------+
`)
fmt.Printf("|%-3s| %-20s   %-4s %-4s %-4s %-4s |\n", "Pos", "Nama Tim", "Pts", "Win", "Lose", "WR%")
for i:= 0 ; i< len(database.DB.Classement); i++{
	fmt.Printf("|%-3d| %-20s   %-4d %-4d %-4d %-4d |\n", i+1,database.DB.Classement[i].Team.Name, database.DB.Classement[i].Pts, database.DB.Classement[i].Win, database.DB.Classement[i].Lose, database.DB.Classement[i].Winrate)
}
	

}

