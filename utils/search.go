package utils

import (
	"esportgacor/database"
)

func linearSearchByID(TeamID int) int{
	idx := -1
	for i := 0; i < len(database.DB.Teams); i++ {
		if database.DB.Teams[i].ID == TeamID {
			idx = i
		}
	}

	return idx
}