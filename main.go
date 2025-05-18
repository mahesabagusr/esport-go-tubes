package main

import (
	"esportgacor/database"
	"esportgacor/handlers"
	"esportgacor/utils"
)

func main() {
	database.InitDB()
	utils.DisplayMainMenu()
	mainSelection := utils.DisplayOption()
	handlers.HandleMainMenu(mainSelection)
}