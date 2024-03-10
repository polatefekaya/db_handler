package config

import (
	"DatabaseHandler/pkg/spesific/player"
)

type AppConfig struct {
	conn string
}

type AppCache struct {
	PlayerModel *player.PlayerRoot
}
