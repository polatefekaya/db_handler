package Players

import (
	pe "DatabaseHandler/pkg/data/entities/Players"
	"time"
)

func (m *Statistic) ToEntity() *pe.StatisticEntity {
	return &pe.StatisticEntity{}
}

func (m *Player) ToEntity() *pe.PlayerEntity {
	return &pe.PlayerEntity{
		Id:           m.Id,
		Name:         m.Name,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		Age:          m.Age,
		Nationality:  m.Nationality,
		Height:       m.Height,
		Weight:       m.Weight,
		Injured:      m.IsInjured,
		PhotoUrl:     m.PhotoUrl,
		BirthDate:    m.Birth.Date.Time,
		BirthCountry: m.Birth.Country,
		BirthPlace:   m.Birth.Place,
		UpdatedAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
	}
}

func (m *League) ToEntity() *pe.LeagueEntity {
	return &pe.LeagueEntity{
		Id:        m.Id,
		Name:      m.Name,
		Country:   m.Country,
		LogoUrl:   m.LogoUrl,
		FlagUrl:   m.FlagUrl,
		Season:    m.Season,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Team) ToEntity() *pe.TeamEntity {
	return &pe.TeamEntity{
		Id:        m.Id,
		Name:      m.Name,
		LogoUrl:   m.LogoUrl,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Card) ToEntity() *pe.CardEntity {
	return &pe.CardEntity{
		Yellow:    m.Yellow,
		YellowRed: m.Yellowred,
		Red:       m.Red,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Dribble) ToEntity() *pe.DribbleEntity {
	return &pe.DribbleEntity{
		Attempts:   m.Attempts,
		Successful: m.Success,
		Past:       m.Past,
		UpdatedAt:  time.Now().UTC(),
		CreatedAt:  time.Now().UTC(),
	}
}

func (m *Duel) ToEntity() *pe.DuelEntity {
	return &pe.DuelEntity{
		Total:     m.Total,
		Won:       m.Total,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Foul) ToEntity() *pe.FoulEntity {
	return &pe.FoulEntity{
		Drawn:     m.Drawn,
		Committed: m.Committed,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Game) ToEntity() *pe.GameEntity {
	return &pe.GameEntity{
		Appearances: m.Appearances,
		Lineups:     m.Lineups,
		Minutes:     m.Minutes,
		Number:      m.Number,
		Position:    m.Position,
		Rating:      m.Rating,
		Captain:     m.IsCaptain,
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
	}
}

func (m *Goal) ToEntity() *pe.GoalEntity {
	return &pe.GoalEntity{
		Total:     m.Total,
		Conceded:  m.Conceded,
		Assists:   m.Assists,
		Saves:     m.Saves,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Pass) ToEntity() *pe.PassEntity {
	return &pe.PassEntity{
		Total:     m.Total,
		Key:       m.Key,
		Accuracy:  m.Accuracy,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Penalty) ToEntity() *pe.PenaltyEntity {
	return &pe.PenaltyEntity{
		Won:       m.Won,
		Committed: m.Committed,
		Scored:    m.Scored,
		Missed:    m.Missed,
		Saved:     m.Saved,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Substitute) ToEntity() *pe.SubstituteEntity {
	return &pe.SubstituteEntity{
		In:        m.In,
		Out:       m.Out,
		Bench:     m.Bench,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}
}

func (m *Tackle) ToEntity() *pe.TackleEntity {
	return &pe.TackleEntity{
		Total:         m.Total,
		Blocks:        m.Blocks,
		Interceptions: m.Interceptions,
		UpdatedAt:     time.Now().UTC(),
		CreatedAt:     time.Now().UTC(),
	}
}
