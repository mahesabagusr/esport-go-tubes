package handlers

import (
	"fmt"
)

func HandleMainMenu(selection int) {
	switch selection {
	case 1:
		HandleTeamMenu()
	default:
		fmt.Println("Invalid selection. Please try again.")
	}
}