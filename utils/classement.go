package utils

import (
	"esportgacor/database"
	"esportgacor/models"
)

func WinratePercentage(wins, losses int) int {
	total := wins + losses
	if total == 0 {
		return 0
	}
	return int(float64(wins) / float64(total) * 100)
}



func WinLoseDraw(TeamName1, TeamName2 string, ScoreT1, ScoreT2 int) {
	var Win1, Lose1, Draw1 int
	var Win2, Lose2, Draw2 int
	if ScoreT1< ScoreT2{
		Lose1= 1
		Win2= 1	
	}else if ScoreT1> ScoreT2 {
		Lose2= 1
		Win1= 1
	}else {
		Draw1= 1
		Draw2= 1
	}
	//Team 1 Processing
	//Is Team 1 In Classement Table Already? If No, Append to Classement
	idx:=linearSearchByTeamNameInClassement(TeamName1)
	if idx == -1 {
		Classement := models.Classement{
		Team: database.DB.Teams[(linearSearchByTeamName(TeamName1))],
		Pts: ((Win1*3)+(Draw1*1)+(Lose1*0)),
		Win: Win1,
		Lose: Lose1,
		Winrate: WinratePercentage(Win1, Lose1),
		}
		database.DB.Classement = append(database.DB.Classement, Classement)
	//if yes then modify the classement instead
	}else {
		database.DB.Classement[idx].Pts+= ((Win1*3)+(Draw1*1)+(Lose1*0))
		database.DB.Classement[idx].Win+= Win1
		database.DB.Classement[idx].Lose+= Lose1
		database.DB.Classement[idx].Winrate= WinratePercentage(database.DB.Classement[idx].Win, database.DB.Classement[idx].Lose)
	}
	//Team 2 Processing
	//Is team 2 in the classement yet? If no then append
	idx=linearSearchByTeamNameInClassement(TeamName2)
	if idx == -1 {
		Classement := models.Classement{
		Team: database.DB.Teams[(linearSearchByTeamName(TeamName2))],
		Pts: ((Win2*3)+(Draw2*1)+(Lose2*0)),
		Win: Win2,
		Lose: Lose2,
		Winrate: WinratePercentage(Win2, Lose2),
		}
		database.DB.Classement = append(database.DB.Classement, Classement)
	//if yes then modify the classement instead
	}else {
		database.DB.Classement[idx].Pts+= ((Win2*3)+(Draw2*1)+(Lose2*0))
		database.DB.Classement[idx].Win+= Win2
		database.DB.Classement[idx].Lose+= Lose2
		database.DB.Classement[idx].Winrate= WinratePercentage(database.DB.Classement[idx].Win, database.DB.Classement[idx].Lose)
	}
	

}