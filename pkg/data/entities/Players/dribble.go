package Players

import "time"

type DribbleEntity struct {
	Id          string
	StatisticId string
	Attempts    int
	Successful  int
	SuccessRate int
	Past        int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
