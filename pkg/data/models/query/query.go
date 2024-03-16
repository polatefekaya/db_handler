package query

type Query struct {
	Players *Players
	Leagues *Leagues
}

func NewQuery() *Query {
	return &Query{}
}

func (m *Query) NewLeagues() *Leagues {
	return &Leagues{}
}
func (m *Query) NewPlayers() *Players {
	return &Players{}
}

func NewSquad() {
	a := players{}
	a.NewPlayer("", "", "", "", "")
}
func zam() {
	a := Query{}
	a.NewPlayers().NewPlayer("", "", "", "", "")
	a.Players.Player.season
}
