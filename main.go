package main

import (
	"esportgacor/database"
	"esportgacor/handlers"
)

func main() {
	database.InitDB()
	handlers.HandleMainMenu()
}