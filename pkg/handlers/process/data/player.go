package data

import (
	e "DatabaseHandler/pkg/data/entities/Players"
	"DatabaseHandler/pkg/data/models/Players"
	"DatabaseHandler/pkg/handlers"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

type PlayerProcessDataHandler struct {
}

type TempStatistic struct {
	statisticId string
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
	*e.ShotEntity
}

func (m *PlayerProcessDataHandler) StartData(root *Players.PlayerRoot) (*e.PlayerEntity, []*e.StatisticEntity, []*TempStatistic) {
	pe := playerData(root)
	se, ts := statisticData(root, pe.Id)
	return pe, se, ts
}

func playerData(root *Players.PlayerRoot) *e.PlayerEntity {
	return root.Responses[0].Player.ToEntity()
}

func teamData(root *Players.PlayerRoot, page int) *e.TeamEntity {
	return root.Responses[0].Statistics[page].Team.ToEntity()
}

func dribbleData(root *Players.PlayerRoot, page int, statId string) *e.DribbleEntity {
	ent := root.Responses[0].Statistics[page].Dribble.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func goalData(root *Players.PlayerRoot, page int, statId string) *e.GoalEntity {
	ent := root.Responses[0].Statistics[page].Goal.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func cardData(root *Players.PlayerRoot, page int, statId string) *e.CardEntity {
	ent := root.Responses[0].Statistics[page].Card.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func duelData(root *Players.PlayerRoot, page int, statId string) *e.DuelEntity {
	ent := root.Responses[0].Statistics[page].Duel.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func foulData(root *Players.PlayerRoot, page int, statId string) *e.FoulEntity {
	ent := root.Responses[0].Statistics[page].Foul.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func gameData(root *Players.PlayerRoot, page int, statId string) *e.GameEntity {
	ent := root.Responses[0].Statistics[page].Game.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func leagueData(root *Players.PlayerRoot, page int) *e.LeagueEntity {
	return root.Responses[0].Statistics[page].League.ToEntity()
}

func passData(root *Players.PlayerRoot, page int, statId string) *e.PassEntity {
	ent := root.Responses[0].Statistics[page].Pass.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func penaltyData(root *Players.PlayerRoot, page int, statId string) *e.PenaltyEntity {
	ent := root.Responses[0].Statistics[page].Penalty.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func statisticData(root *Players.PlayerRoot, playerId int) ([]*e.StatisticEntity, []*TempStatistic) {
	es, err := entitiesProcess(root)
	if err != nil {
		log.Fatal(err)

	}

	sel := make([]*e.StatisticEntity, 0, cap(es))

	for i := range len(es) {
		sel = append(sel, makeStatEntity(es[i], playerId))
	}

	return sel, es
}

func entitiesProcess(root *Players.PlayerRoot) ([]*TempStatistic, error) {
	sl := len(root.Responses[0].Statistics)
	es := make([]*TempStatistic, 0, sl)
	for i := range sl {
		es = append(es, processBuffer(root, i))
	}

	if len(es) != sl {
		return nil, errors.New(fmt.Sprintf("Entities Process couldn't process everything. Expected Len: %d, Got: %d", sl, len(es)))
	}
	return es, nil
}

func makeStatEntity(temp *TempStatistic, playerId int) *e.StatisticEntity {
	se := e.StatisticEntity{
		Id:           temp.statisticId,
		PlayerId:     playerId,
		TeamId:       temp.TeamEntity.Id,
		LeagueId:     temp.LeagueEntity.Id,
		GameId:       temp.GameEntity.Id,
		SubstituteId: temp.SubstituteEntity.Id,
		ShotId:       temp.ShotEntity.Id,
		GoalId:       temp.GoalEntity.Id,
		PassId:       temp.PassEntity.Id,
		TackleId:     temp.TackleEntity.Id,
		DuelId:       temp.DuelEntity.Id,
		DribbleId:    temp.DribbleEntity.Id,
		FoulId:       temp.FoulEntity.Id,
		CardId:       temp.CardEntity.Id,
		PenaltyId:    temp.PenaltyEntity.Id,
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
	}
	return &se
}

func processBuffer(root *Players.PlayerRoot, i int) *TempStatistic {
	statisticId := uuid.New().String()
	seo := TempStatistic{
		statisticId:      statisticId,
		TeamEntity:       teamData(root, i),
		LeagueEntity:     leagueData(root, i),
		DribbleEntity:    dribbleData(root, i, statisticId),
		GoalEntity:       goalData(root, i, statisticId),
		CardEntity:       cardData(root, i, statisticId),
		DuelEntity:       duelData(root, i, statisticId),
		FoulEntity:       foulData(root, i, statisticId),
		GameEntity:       gameData(root, i, statisticId),
		PassEntity:       passData(root, i, statisticId),
		PenaltyEntity:    penaltyData(root, i, statisticId),
		SubstituteEntity: substituteData(root, i, statisticId),
		TackleEntity:     tackleData(root, i, statisticId),
		ShotEntity:       shotData(root, i, statisticId),
	}
	return &seo
}

func shotData(root *Players.PlayerRoot, page int, statId string) *e.ShotEntity {
	ent := root.Responses[0].Statistics[page].Shot.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func substituteData(root *Players.PlayerRoot, page int, statId string) *e.SubstituteEntity {
	ent := root.Responses[0].Statistics[page].Substitute.ToEntity()
	return handlers.FillStatId(ent, statId)
}

func tackleData(root *Players.PlayerRoot, page int, statId string) *e.TackleEntity {
	ent := root.Responses[0].Statistics[page].Tackle.ToEntity()
	return handlers.FillStatId(ent, statId)
}
