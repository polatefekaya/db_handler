package pkg

import (
	"DatabaseHandler/pkg/config"
)

var c Cache

// Cache player
type Cache struct {
	ac *config.AppCache
}

func CachePlayer(ac *config.AppCache) *Cache {
	return &Cache{
		ac: ac,
	}
}
