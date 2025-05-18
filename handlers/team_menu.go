package handlers

import (
	"esportgacor/utils"
)

func HandleTeamMenu() {
	utils.DisplayManageTeamMenu()
	selection := utils.DisplayOption()

	switch selection {
	case 1:
		utils.AddTeam()
		HandleTeamMenu()
	case 2:
		utils.Modifyteam()
		HandleTeamMenu()
	case 3:
		utils.DeleteTeam()
		HandleTeamMenu()
	case 4 : 
	 	HandleMainMenu()
	default:
		println("Invalid selection. Please try again.")
	}
}