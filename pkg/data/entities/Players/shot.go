package Players

import "time"

type ShotEntity struct {
	Id        string
	Total     int
	On        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
