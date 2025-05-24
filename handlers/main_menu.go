package handlers

import (
	"esportgacor/utils"
	"fmt"
)

func HandleMainMenu() {
	for{
		utils.DisplayMainMenu()
		selection := utils.ScanNumber("Pilih Opsi: ")
	
	switch selection {
	case 1:
		HandleTeamMenu()
	case 2:
		utils.DisplayAllTeamsAndPlayersMenu()
	case 4:
		HandleTeamMenu()
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}