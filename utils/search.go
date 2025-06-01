package utils

import (
	"esportgacor/database"
)

func linearSearchByTeamID(TeamID int) int{
	idx := -1
	for i := 0; i < len(database.DB.Teams); i++ {
		if database.DB.Teams[i].ID == TeamID {
			idx = i
		}
	}

	return idx
}

func binarySearchByMatchID(MatchID int) int {
	var left, right, mid int
	idx := -1 // Important: initialize to -1
	left = 0
	right = len(database.DB.Matches) - 1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if database.DB.Matches[mid].MatchID == MatchID {
			idx = mid
		} else if database.DB.Matches[mid].MatchID < MatchID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return idx
}

func linearSearchByTeamNameInClassement(TeamName string) int{
	idx := -1
	for i := 0; i < len(database.DB.Classement); i++ {
		if database.DB.Classement[i].Team.Name == TeamName {
			idx = i
		}
	}

	return idx
}

func linearSearchByTeamName(TeamName string) int{
	idx := -1
	for i := 0; i < len(database.DB.Teams); i++ {
		if database.DB.Teams[i].Name == TeamName {
			idx = i
		}
	}

	return idx
}