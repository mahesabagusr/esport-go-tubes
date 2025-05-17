package main

import (
	"esportgacor/handlers"
	"esportgacor/utils"
)

func main() {
	utils.DisplayMainMenu()
	mainSelection := utils.DisplayOption()
	handlers.HandleMainMenu(mainSelection)
}