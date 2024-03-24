package data

import (
	e "DatabaseHandler/pkg/data/entities/Players"
	"DatabaseHandler/pkg/data/models/Players"
	proc "DatabaseHandler/pkg/interfaces/process"
	"errors"
	"fmt"
	"log"
)

type PlayerProcessDataHandler struct {
}

type IPlayerDataProcess interface {
	StartData(root *Players.PlayerRoot)
	proc.IPlayerProcess
}

type tempEntities struct {
	*e.TeamEntity
	*e.LeagueEntity
	*e.DribbleEntity
	*e.GoalEntity
	*e.CardEntity
	*e.DuelEntity
	*e.FoulEntity
	*e.GameEntity
	*e.PassEntity
	*e.PenaltyEntity
	*e.SubstituteEntity
	*e.TackleEntity
}

func (m *PlayerProcessDataHandler) StartData(root *Players.PlayerRoot) {
	playerData(root)
	teamData(root, 0)
	dribbleData(root, 0)
	statisticData(root)
}

func (m *PlayerProcessDataHandler) PlayerProcess() {
}

func (m *PlayerProcessDataHandler) LeagueProcess() {

}

func (m *PlayerProcessDataHandler) TeamProcess() {

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
	es, err := entitiesProcess(root)
	if err != nil {
		log.Fatal(err)
	}
	//uuid.New()

}

func entitiesProcess(root *Players.PlayerRoot) ([]*tempEntities, error) {
	sl := len(root.Responses[0].Statistics)
	es := make([]*tempEntities, 0, sl)
	for i := range sl {
		es = append(es, processBuffer(root, i))
	}

	if len(es) != sl {
		return nil, errors.New(fmt.Sprintf("Entities Process couldn't process everything. Expected Len: %d, Got: %d", sl, len(es)))
	}
	return es, nil
}

func processBuffer(root *Players.PlayerRoot, i int) *tempEntities {
	seo := tempEntities{
		TeamEntity:       teamData(root, i),
		LeagueEntity:     leagueData(root, i),
		DribbleEntity:    dribbleData(root, i),
		GoalEntity:       goalData(root, i),
		CardEntity:       cardData(root, i),
		DuelEntity:       duelData(root, i),
		FoulEntity:       foulData(root, i),
		GameEntity:       gameData(root, i),
		PassEntity:       passData(root, i),
		PenaltyEntity:    penaltyData(root, i),
		SubstituteEntity: substituteData(root, i),
		TackleEntity:     tackleData(root, i),
	}
	return &seo
}

func substituteData(root *Players.PlayerRoot, page int) *e.SubstituteEntity {
	return root.Responses[0].Statistics[page].Substitute.ToEntity()
}

func tackleData(root *Players.PlayerRoot, page int) *e.TackleEntity {
	return root.Responses[0].Statistics[page].Tackle.ToEntity()
}
