package memory

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	c *cache.Cache
)

func init() {
	c = cache.New(cache.NoExpiration, 5*time.Minute)
}

func Cache() *cache.Cache {
	return c
}
