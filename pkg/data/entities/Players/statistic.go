package Players

import "time"

type StatisticEntity struct {
	Id           string
	PlayerId     int
	TeamId       int
	LeagueId     int
	GameId       string
	SubstituteId string
	ShotId       string
	GoalId       string
	PassId       string
	TackleId     string
	DuelId       string
	DribbleId    string
	FoulId       string
	CardId       string
	PenaltyId    string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
