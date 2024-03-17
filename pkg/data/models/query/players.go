package query

import (
	"errors"
	"strings"
)

type Players struct {
	Player *Player
	Squad  *Squad
}

type IPlayers interface {
	Generate(string)
}

type Player struct {
	PlayerId string
	Base     string "players"
	Season   string
	Search   string
}

type Squad struct {
	TeamId string
}

func (m *Player) Generate(playerId, season, search string) *Player {
	m.PlayerId = playerId
	m.Season = season
	m.Search = search

	return m
}

func (m *Player) Build() (string, error) {
	sb := strings.Builder{}
	sb.WriteByte('?')
	sb.WriteString(m.Base)

	if strings.Compare(m.PlayerId, "") == 0 {
		return "", errors.New("no playerId value provided for query")
	}
	sb.WriteString("id=")
	sb.WriteString(m.PlayerId)
	sb.WriteByte('&')

	if strings.Compare(m.Season, "") == 0 {
		return "", errors.New("no season value provided for query")
	}
	sb.WriteString("season=")
	sb.WriteString(m.Season)

	if strings.Compare(m.Search, "") != 0 {
		sb.WriteByte('&')
		sb.WriteString("search=")
		sb.WriteString(m.Search)
	}
	sb.WriteByte('/')

	return sb.String(), nil
}

func (m *Squad) Generate(teamId string) {
	m.TeamId = teamId
}
