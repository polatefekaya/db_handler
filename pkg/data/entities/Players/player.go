package Players

import "time"

type PlayerEntity struct {
	Id           int
	Name         string
	FirstName    string
	LastName     string
	Age          int
	Nationality  string
	Height       string
	Weight       string
	Injured      bool
	PhotoUrl     string
	BirthDate    time.Time
	BirthPlace   string
	BirthCountry string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
