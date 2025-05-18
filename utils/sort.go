package utils

import (
	"esportgacor/models"
)
func SortTeamsWithName(teams []models.Team) []models.Team {
	sorted := make([]models.Team, len(teams))
	copy(sorted, teams)

	// Insertion sort
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