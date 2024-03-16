package Players

import "time"

type TeamEntity struct {
	Id        int
	Name      string
	LogoUrl   string
	UpdatedAt time.Time
	CreatedAt time.Time
}
