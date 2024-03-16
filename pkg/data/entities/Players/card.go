package Players

import "time"

type CardEntity struct {
	Id          string
	StatisticId string
	Yellow      int
	YellowRed   int
	Red         int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
