package query

type Players struct {
	Player *Player
	Squad  *Squad
}

type IPlayers interface {
	Generate(string)
}

type Player struct {
	PlayerId string
	TeamId   string
	LeagueId string
	Season   string
	Search   string
}

type Squad struct {
	TeamId string
}

func (m *Player) Generate(playerId, teamId, leagueId, season, search *string) {
	m.PlayerId = *playerId
	m.TeamId = *teamId
	m.LeagueId = *leagueId
	m.Season = *season
	m.Search = *search
}

func (m *Squad) Generate(teamId string) {
	m.TeamId = teamId
}
