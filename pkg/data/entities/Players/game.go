package Players

import "time"

type GameEntity struct {
	Id          string
	StatisticId string
	Appearances int
	Lineups     int
	Minutes     int
	Number      int
	Position    string
	Rating      string
	Captain     bool
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
