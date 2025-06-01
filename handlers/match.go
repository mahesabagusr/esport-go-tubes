package handlers

import "esportgacor/utils"

func HandleMatchMenu() {
	for {
	utils.DisplayManageMatchMenu()
	selection := utils.ScanNumber("Pilih Opsi: ")
	switch selection {
	case 1:
		utils.AddMatches()
		return
	case 2:
		utils.EditMatches()
		return
	case 3:
		utils.ViewMatches()
		utils.ScanString("Tekan Enter untuk melanjutkan...")
		return
	case 4 : 
	 	return  
	default:
		println("Invalid selection. Please try again.")
	}
	}
}