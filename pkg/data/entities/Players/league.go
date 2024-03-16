package Players

import "time"

type LeagueEntity struct {
	Id        int
	Name      string
	Country   string
	LogoUrl   string
	FlagUrl   string
	Season    int
	UpdatedAt time.Time
	CreatedAt time.Time
}
