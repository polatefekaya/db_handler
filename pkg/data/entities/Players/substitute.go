package Players

import "time"

type SubstituteEntity struct {
	Id          string
	StatisticId string
	In          int
	Out         int
	Bench       int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
