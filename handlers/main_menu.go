package handlers

import (
	"esportgacor/utils"
	"fmt"
)

func HandleMainMenu() {
	utils.DisplayMainMenu()
	selection := utils.DisplayOption()
	
	switch selection {
	case 1:
		HandleTeamMenu()
	default:
		fmt.Println("Invalid selection. Please try again.")
	}
}