package player

type PlayerRoot struct {
	Responses []Response `json:"response"`
}

type Response struct {
	Player     Player      `json:"player"`
	Statistics []Statistic `json:"statistics"`
}

type Player struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Age         int    `json:"age"`
	Birth       Birth  `json:"birth"`
	Nationality string `json:"nationality"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	IsInjured   bool   `json:"injured"`
	PhotoUrl    string `json:"photo"`
}

type Birth struct {
	Date    string `json:"date"`
	Place   string `json:"place"`
	Country string `json:"country"`
}

type Statistic struct {
	Team       Team       `json:"team"`
	League     League     `json:"league"`
	Game       Game       `json:"games"`
	Substitute Substitute `json:"substitutes"`
	Shot       Shot       `json:"shots"`
	Goal       Goal       `json:"goals"`
	Pass       Pass       `json:"passes"`
	Tackle     Tackle     `json:"tackles"`
	Duel       Duel       `json:"duels"`
	Dribble    Dribble    `json:"dribbles"`
	Foul       Foul       `json:"fouls"`
	Card       Card       `json:"cards"`
	Penalty    Penalty    `json:"penalty"`
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
