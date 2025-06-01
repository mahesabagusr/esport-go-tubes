package handlers

import (
	"esportgacor/utils"
)

func HandleTeamMenu() {
	for {
	utils.DisplayManageTeamMenu()
	selection := utils.ScanNumber("Pilih Opsi: ")

	switch selection {
	case 1:
		utils.AddTeam()
		return
	case 2:
		utils.ModifyTeam()
		return
	case 3:
		utils.DeleteTeam()
		return
	case 4 : 
	 	return  
	default:
		println("Invalid selection. Please try again.")
	}
	}
}