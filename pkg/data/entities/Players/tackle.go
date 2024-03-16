package Players

import "time"

type TackleEntity struct {
	Id            string
	StatisticId   string
	Total         int
	Blocks        int
	Interceptions int
	UpdatedAt     time.Time
	CreatedAt     time.Time
}
