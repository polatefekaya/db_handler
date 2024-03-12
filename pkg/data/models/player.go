package models

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
	Id      int    `json:"id"`
	Name    string `json:"name"`
	LogoUrl string `json:"logo"`
}

type League struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	LogoUrl string `json:"logo"`
	FlagUrl string `json:"flag"`
	Season  int    `json:"season"`
}

type Game struct {
	Appearances int    `json:"appearances"`
	Lineups     int    `json:"lineups"`
	Minutes     int    `json:"minutes"`
	Number      int    `json:"number"`
	Position    string `json:"position"`
	Rating      string `json:"rating"`
	IsCaptain   bool   `json:"captain"`
}

type Substitute struct {
	In    int `json:"in"`
	Out   int `json:"out"`
	Bench int `json:"bench"`
}

type Shot struct {
	Total int `json:"total"`
	On    int `json:"on"`
}

type Goal struct {
	Total    int `json:"total"`
	Conceded int `json:"conceded"`
	Assists  int `json:"assists"`
	Saves    int `json:"saves"`
}

type Pass struct {
	Total    int `json:"total"`
	Key      int `json:"key"`
	Accuracy int `json:"accuracy"`
}

type Tackle struct {
	Total         int `json:"total"`
	Blocks        int `json:"blocks"`
	Interceptions int `json:"interceptions"`
}

type Duel struct {
	Total int `json:"total"`
	Won   int `json:"won"`
}

type Dribble struct {
	Attempts int `json:"attempts"`
	Success  int `json:"success"`
	Past     int `json:"past"`
}

type Foul struct {
	Drawn     int `json:"drawn"`
	Committed int `json:"committed"`
}

type Card struct {
	Yellow    int `json:"yellow"`
	Yellowred int `json:"yellowred"`
	Red       int `json:"red"`
}

type Penalty struct {
	Won       int `json:"won"`
	Committed int `json:"committed"`
	Scored    int `json:"scored"`
	Missed    int `json:"missed"`
	Saved     int `json:"saved"`
}
