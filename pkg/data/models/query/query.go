package query

type Query struct {
	Players *Players
	Leagues *Leagues
}

type Players struct {
	Squad  *squad
	Player *player
}

type Leagues struct {
	League league
}

type league struct {
}

type player struct {
	id       string
	teamId   string
	leagueId string
	season   string
	search   string
}

type squad struct {
}

func (m *Query) NewLeagues() *Leagues {
	return &Leagues{}
}
func (m *Query) NewPlayers() *Players {
	return &Players{}
}

func (m *Players) NewPlayer(playerId, teamId, leagueId, season, search string) {
	m.Player = &player{
		id:       playerId,
		teamId:   teamId,
		leagueId: leagueId,
		season:   season,
		search:   search,
	}
}

func NewSquad() {
	a := Players{}
	a.NewPlayer("", "", "", "", "")
}
func zam() {
	a := Query{}
	a.NewPlayers().NewPlayer("", "", "", "", "")
	a.Players.Player.season
}
