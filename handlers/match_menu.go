package handlers

import "esportgacor/utils"

func HandleMatchMenu() {
	for {
	utils.DisplayManageTeamMenu()
	selection := utils.ScanNumber("Pilih Opsi: ")
	switch selection {
	case 1:
		utils.ViewMatches()
		return
	case 2:
		//utils.AddMatches()
		return
	case 3:
		//utils.EditMatches()
		return
	case 4 : 
	 	return  
	default:
		println("Invalid selection. Please try again.")
	}
	}
}