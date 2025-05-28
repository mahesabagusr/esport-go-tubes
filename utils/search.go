package utils

import (
	"esportgacor/database"
	"fmt"
)

func linearSearch(TeamID int){
	idx := -1
	for i := 0; i < len(database.DB.Teams); i++ {
		if database.DB.Teams[i].ID == TeamID {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Tim tidak ditemukan!")
		return
	}
}