package entities

import "time"

type PlayerEntity struct {
	Id           int
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
