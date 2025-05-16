package main

import (
	"esportgacor/utils"
	"fmt"
)

func main(){
	// Main Interface
	var mainSelection, teamSelection int
	utils.DisplayMainMenu()
	fmt.Scanln(&mainSelection)
	switch mainSelection{
	case 1 :
		utils.DisplayManageTeamMenu()
		fmt.Scanln(&teamSelection)
		switch teamSelection{
		case 1:
			utils.Addteam()
		case 2:
			utils.Modifyteam()
		case 3:
			utils.DeleteTeam()
		}
	case 2 :
		fmt.Println("Belum Dibuat Cik")
	}


}

