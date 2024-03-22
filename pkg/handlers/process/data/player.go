package data

import (
	e "DatabaseHandler/pkg/data/entities/Players"
	"DatabaseHandler/pkg/data/models/Players"
)

type PlayerProcessDataHandler struct {
}

func (m *PlayerProcessDataHandler) StartData(root *Players.PlayerRoot) {
	playerData(root)
	teamData(root, 0)
	dribbleData(root, 0)
	statisticData(root)
}

func playerData(root *Players.PlayerRoot) *e.PlayerEntity {
	return root.Responses[0].Player.ToEntity()
}

func teamData(root *Players.PlayerRoot, page int) *e.TeamEntity {
	return root.Responses[0].Statistics[page].Team.ToEntity()
}

func dribbleData(root *Players.PlayerRoot, page int) *e.DribbleEntity {
	return root.Responses[0].Statistics[page].Dribble.ToEntity()
}

func goalData(root *Players.PlayerRoot, page int) *e.GoalEntity {
	return nil
}

func cardData(root *Players.PlayerRoot, page int) *e.CardEntity {
	return nil
}

func duelData(root *Players.PlayerRoot, page int) {

}

func foulData(root *Players.PlayerRoot, page int) {

}

func gameData(root *Players.PlayerRoot, page int) {

}

func leagueData(root *Players.PlayerRoot, page int) {

}

func passData(root *Players.PlayerRoot, page int) {

}

func penaltyData(root *Players.PlayerRoot, page int) {

}

func statisticData(root *Players.PlayerRoot) {
	var sd e.StatisticEntity
	sd.TeamId = 12
}

func substituteData(root *Players.PlayerRoot, page int) {

}

func tackleData(root *Players.PlayerRoot, page int) {

}
