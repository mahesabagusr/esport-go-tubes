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
	case 2:
		utils.Modifyteam()
	case 3:
		utils.DeleteTeam()
	default:
		println("Invalid selection. Please try again.")
	}
}