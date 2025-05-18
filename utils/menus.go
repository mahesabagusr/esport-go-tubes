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
+-----------------------------------------------------------+
`)
}

func DisplayManageTeamMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                       MANAGE TEAM                          |
+-----------------------------------------------------------+
| 1. ADD TEAM                                               |
| 2. MODIFY TEAM                                            |
| 3. DELETE TEAM                                            |
| 4. BACK                                            				|
+-----------------------------------------------------------+
`)
}

func DisplayAllTeamsMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                   DAFTAR SEMUA TIM & PEMAIN              |
+-----------------------------------------------------------+
`)
	for _, team := range database.DB.Teams {
		fmt.Printf("| ID: %-3d Nama Tim: %-20s Coach: %-10s |\n", team.ID, team.Name, team.Coach)
		fmt.Println("| Pemain:")
		for _, player := range team.Players {
			fmt.Printf("|   - ID: %-3d Nama: %-15s Kills: %-3d Deaths: %-3d |\n", player.ID, player.Name, player.Kills, player.Deaths)
		}
		fmt.Println("+-----------------------------------------------------------+")
	}
}
