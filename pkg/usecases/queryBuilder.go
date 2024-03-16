package usecases

import "DatabaseHandler/pkg/data/models/query"

type QueryBuilder struct {
	Query *query.Query
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (m *QueryBuilder) Build(player *query.Player) string {
	qs := player.Generate(",", ",", ",", ",", ",")

}
