package player

import "DatabaseHandler/pkg/data/models"

type Types interface {
	models.PlayerRoot | models.Player | models.Team | models.League | models.Statistic | models.Goal | models.Shot | models.Game | models.Birth | models.Substitute | models.Pass | models.Tackle | models.Duel | models.Dribble | models.Foul | models.Card | models.Penalty
}
