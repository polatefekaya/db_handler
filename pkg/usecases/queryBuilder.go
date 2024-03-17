package usecases

import (
	"DatabaseHandler/pkg/data/models/query"
	"log"
)

type QueryBuilder struct {
	Query *query.Query
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (m *QueryBuilder) Build(player *query.Player) string {
	q, err := player.Generate(",", ",", ",").Build()
	if err != nil {
		log.Fatal(err)
	}
	return q
}
