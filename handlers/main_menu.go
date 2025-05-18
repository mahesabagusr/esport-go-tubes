package handlers

import (
	"esportgacor/utils"
	"fmt"
)

func HandleMainMenu() {
	for{
		utils.DisplayMainMenu()
		selection := utils.ScanNumber()
	
	switch selection {
	case 1:
		HandleTeamMenu()
	case 2:
		utils.DisplayAllTeamsMenu()
	default:
		fmt.Println("Invalid selection. Please try again.")
		}
	}
}