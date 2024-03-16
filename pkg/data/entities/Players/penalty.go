package Players

import "time"

type PenaltyEntity struct {
	Id          string
	StatisticId string
	Won         int
	Committed   int
	Scored      int
	Missed      int
	ScoreRate   int
	Saved       int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
