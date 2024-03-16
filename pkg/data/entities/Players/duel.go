package Players

import "time"

type DuelEntity struct {
	Id          string
	StatisticId string
	Total       int
	Won         int
	Success     int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
