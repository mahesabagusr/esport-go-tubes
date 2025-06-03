package utils

import (
	"esportgacor/models"
)
func SortTeamsWithName(teams []models.Team) []models.Team {
	sorted := make([]models.Team, len(teams))
	copy(sorted, teams)

	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j].Name > key.Name {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}

	return sorted
}

func SortTeamsByPts(classements []models.Classement) {
	n := len(classements)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if classements[j].Pts > classements[maxIdx].Pts {
				maxIdx = j
			}
		}
		classements[i], classements[maxIdx] = classements[maxIdx], classements[i]
	}
}


