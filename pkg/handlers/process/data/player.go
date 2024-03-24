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
	return root.Responses[0].Statistics[page].Goal.ToEntity()
}

func cardData(root *Players.PlayerRoot, page int) *e.CardEntity {
	return root.Responses[0].Statistics[page].Card.ToEntity()
}

func duelData(root *Players.PlayerRoot, page int) *e.DuelEntity {
	return root.Responses[0].Statistics[page].Duel.ToEntity()
}

func foulData(root *Players.PlayerRoot, page int) *e.FoulEntity {
	return root.Responses[0].Statistics[page].Foul.ToEntity()
}

func gameData(root *Players.PlayerRoot, page int) *e.GameEntity {
	return root.Responses[0].Statistics[page].Game.ToEntity()
}

func leagueData(root *Players.PlayerRoot, page int) *e.LeagueEntity {
	return root.Responses[0].Statistics[page].League.ToEntity()
}

func passData(root *Players.PlayerRoot, page int) *e.PassEntity {
	return root.Responses[0].Statistics[page].Pass.ToEntity()
}

func penaltyData(root *Players.PlayerRoot, page int) *e.PenaltyEntity {
	return root.Responses[0].Statistics[page].Penalty.ToEntity()
}

func statisticData(root *Players.PlayerRoot) {
	var sd e.StatisticEntity
	sl := len(root.Responses[0].Statistics)
	for i := range sl {
		processBuffer(root, i)
	}
	sd.TeamId = 12
}

func processBuffer(root *Players.PlayerRoot, i int) {
	teamData(root, i)
	leagueData(root, i)
	dribbleData(root, i)
	goalData(root, i)
	cardData(root, i)
	duelData(root, i)
	foulData(root, i)
	gameData(root, i)
	passData(root, i)
	penaltyData(root, i)
	substituteData(root, i)
	tackleData(root, i)
}

func substituteData(root *Players.PlayerRoot, page int) *e.SubstituteEntity {
	return root.Responses[0].Statistics[page].Substitute.ToEntity()
}

func tackleData(root *Players.PlayerRoot, page int) *e.TackleEntity {
	return root.Responses[0].Statistics[page].Tackle.ToEntity()
}
