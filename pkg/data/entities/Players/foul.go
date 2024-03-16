package Players

import "time"

type FoulEntity struct {
	Id          string
	StatisticId string
	Drawn       int
	Committed   int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
