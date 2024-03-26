package data

import (
	"DatabaseHandler/internals"
	log2 "DatabaseHandler/internals/log"
	e "DatabaseHandler/pkg/data/entities/Players"
	"DatabaseHandler/pkg/data/models/Players"
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
	pe := root.Responses[0].Player.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", pe)
	return pe
}

func teamData(root *Players.PlayerRoot, page int) *e.TeamEntity {
	te := root.Responses[0].Statistics[page].Team.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", te)
	return te
}

func leagueData(root *Players.PlayerRoot, page int) *e.LeagueEntity {
	le := root.Responses[0].Statistics[page].League.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", le)
	return le
}

func dribbleData(root *Players.PlayerRoot, page int, statId string) *e.DribbleEntity {
	ent := root.Responses[0].Statistics[page].Dribble.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func goalData(root *Players.PlayerRoot, page int, statId string) *e.GoalEntity {
	ent := root.Responses[0].Statistics[page].Goal.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func cardData(root *Players.PlayerRoot, page int, statId string) *e.CardEntity {
	ent := root.Responses[0].Statistics[page].Card.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func duelData(root *Players.PlayerRoot, page int, statId string) *e.DuelEntity {
	ent := root.Responses[0].Statistics[page].Duel.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func foulData(root *Players.PlayerRoot, page int, statId string) *e.FoulEntity {
	ent := root.Responses[0].Statistics[page].Foul.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func gameData(root *Players.PlayerRoot, page int, statId string) *e.GameEntity {
	ent := root.Responses[0].Statistics[page].Game.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func passData(root *Players.PlayerRoot, page int, statId string) *e.PassEntity {
	ent := root.Responses[0].Statistics[page].Pass.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func penaltyData(root *Players.PlayerRoot, page int, statId string) *e.PenaltyEntity {
	ent := root.Responses[0].Statistics[page].Penalty.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func statisticData(root *Players.PlayerRoot, playerId int) ([]*e.StatisticEntity, []*TempStatistic) {
	log2.INFO("Statistic data process started.")
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
	log2.INFO("Statistic temporary slice generated.", "Capacity: ", sl)
	for i := range sl {
		es = append(es, processBuffer(root, i))
	}

	if len(es) != sl {
		return nil, errors.New(fmt.Sprintf("Entities Process couldn't process everything. Expected Len: %d, Got: %d", sl, len(es)))
	}
	log2.INFO("Statistic temporary slice filled")
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
	log2.DEBUG("Processing model to entity ", "Entity: ", se)
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
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func substituteData(root *Players.PlayerRoot, page int, statId string) *e.SubstituteEntity {
	ent := root.Responses[0].Statistics[page].Substitute.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}

func tackleData(root *Players.PlayerRoot, page int, statId string) *e.TackleEntity {
	ent := root.Responses[0].Statistics[page].Tackle.ToEntity()
	log2.DEBUG("Processing model to entity ", "Entity: ", ent)
	return internals.FillStatId(ent, statId)
}
