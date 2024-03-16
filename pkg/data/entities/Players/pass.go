package Players

import "time"

type PassEntity struct {
	Id          string
	StatisticId string
	Total       int
	Key         int
	Accuracy    int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
