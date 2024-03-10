package models

import (
	"time"
)

type RootModel struct {
	responses []responseModel
}

type responseModel struct {
	player     playerModel
	statistics []statisticModel
}

type playerModel struct {
	firstName   string     `json:"firstname"`
	lastName    string     `json:"lastname"`
	age         int        `json:"age"`
	birth       birthModel `json:"birth"`
	nationality string     `json:"nationality"`
	height      string     `json:"height"`
	weight      string     `json:"weight"`
	isInjured   bool       `json:"injured"`
	photoUrl    string     `json:"photo"`
}

type birthModel struct {
	date    time.Time `json:"date"`
	place   string    `json:"place"`
	country string    `json:"country"`
}

type statisticModel struct {
	team   teamModel
	leauge leagueModel
}

type teamModel struct {
	id      int    `json:"id"`
	name    string `json:"name"`
	logoUrl string `json:"logo"`
}

type leagueModel struct {
	id      int    `json:"id"`
	name    int    `json:"name"`
	country string `json:"country"`
	logoUrl string `json:"logo"`
	flagUrl string `json:"flag"`
	season  int    `json:"season"`
}
