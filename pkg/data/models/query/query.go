package query

type Query struct {
	Players *Players
	Leagues *Leagues
}

func NewQuery() *Query {
	return &Query{}
}

func NewLeague() *Leagues {
	return &Leagues{}
}
func NewPlayer() *Player {
	return &Player{}
}
