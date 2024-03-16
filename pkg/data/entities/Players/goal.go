package Players

import "time"

type GoalEntity struct {
	Id          string
	StatisticId string
	Total       int
	Conceded    int
	Assists     int
	Saves       int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
