package player

import (
	"time"
)

type PlayerRoot struct {
	responses []Response `json:"response"`
}

type Response struct {
	player     Player      `json:"player"`
	statistics []Statistic `json:"statistics"`
}

type Player struct {
	firstName   string `json:"firstname"`
	lastName    string `json:"lastname"`
	age         int    `json:"age"`
	birth       Birth  `json:"birth"`
	nationality string `json:"nationality"`
	height      string `json:"height"`
	weight      string `json:"weight"`
	isInjured   bool   `json:"injured"`
	photoUrl    string `json:"photo"`
}

type Birth struct {
	date    time.Time `json:"date"`
	place   string    `json:"place"`
	country string    `json:"country"`
}

type Statistic struct {
	team       Team       `json:"team"`
	leauge     League     `json:"leauge"`
	game       Game       `json:"games"`
	substitute Substitute `json:"substitutes"`
	shot       Shot       `json:"shots"`
	goal       Goal       `json:"goals"`
	pass       Pass       `json:"passes"`
	tackle     Tackle     `json:"tackles"`
	duel       Duel       `json:"duels"`
	dribble    Dribble    `json:"dribbles"`
	foul       Foul       `json:"fouls"`
	card       Card       `json:"cards"`
	penalty    Penalty    `json:"penalty"`
}

type Team struct {
	id      int    `json:"id"`
	name    string `json:"name"`
	logoUrl string `json:"logo"`
}

type League struct {
	id      int    `json:"id"`
	name    int    `json:"name"`
	country string `json:"country"`
	logoUrl string `json:"logo"`
	flagUrl string `json:"flag"`
	season  int    `json:"season"`
}

type Game struct {
	appearences int    `json:"appearences"`
	lineups     int    `json:"lineups"`
	minutes     int    `json:"minutes"`
	number      int    `json:"number"`
	position    string `json:"position"`
	rating      string `json:"rating"`
	isCaptain   bool   `json:"captain"`
}

type Substitute struct {
	in    int `json:"in"`
	out   int `json:"out"`
	bench int `json:"bench"`
}

type Shot struct {
	total int `json:"total"`
	on    int `json:"on"`
}

type Goal struct {
	total    int `json:"total"`
	conceded int `json:"conceded"`
	assists  int `json:"assists"`
	saves    int `json:"saves"`
}

type Pass struct {
	total    int `json:"total"`
	key      int `json:"key"`
	accuracy int `json:"accuracy"`
}

type Tackle struct {
	total         int `json:"total"`
	blocks        int `json:"blocks"`
	interceptions int `json:"interceptions"`
}

type Duel struct {
	total int `json:"total"`
	won   int `json:"won"`
}

type Dribble struct {
	attempts int `json:"attempts"`
	success  int `json:"success"`
	past     int `json:"past"`
}

type Foul struct {
	drawn     int `json:"drawn"`
	committed int `json:"committed"`
}

type Card struct {
	yellow    int `json:"yellow"`
	yellowred int `json:"yellowred"`
	red       int `json:"red"`
}

type Penalty struct {
	won      int `json:"won"`
	commited int `json:"commited"`
	scored   int `json:"scored"`
	missed   int `json:"missed"`
	saved    int `json:"saved"`
}
