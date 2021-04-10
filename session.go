// @author mr.long

package wineglass

import (
	"encoding/gob"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	DefaultName = "session"

	// SessionCookieStore Session Store use cookie
	SessionCookieStore = "cookie"
	// SessionRedisStore Session Store use redis
	SessionRedisStore = "redis"
)

func (m *Middleware) Sessions() gin.HandlerFunc {
	conf := m.Session
	var sessionStore redis.Store

	switch conf.Store {
	case SessionCookieStore:
		sessionStore = cookie.NewStore([]byte(conf.Secret))
	case SessionRedisStore:
		var err error
		sessionStore, err = redis.NewStoreWithDB(
			conf.Redis.Size,
			"tcp",
			conf.Redis.Address,
			conf.Redis.Password,
			conf.Redis.DB,
			[]byte(conf.Secret),
		)
		if err != nil {
			log.Fatalf("%+v\n", errors.New(err.Error()))
		}
	default:
		sessionStore = cookie.NewStore([]byte(conf.Secret))
	}

	sessionStore.Options(*conf.Options)

	for _, val := range conf.GobModels {
		gob.Register(val)
	}

	if len(conf.Name) == 0 {
		conf.Name = DefaultName
	}
	return sessions.Sessions(conf.Name, sessionStore)
}
