package middleware

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/consts"
)

const (
	SessionName = "session"
)

type Session struct{}

func (s *Session) Init() ([]gin.HandlerFunc, error) {
	var sessionStore redis.Store

	switch config.Session.Store {
	case consts.CookieStore:
		sessionStore = cookie.NewStore([]byte(config.Session.Secret))
	case consts.MemoryStore:
		sessionStore = memstore.NewStore([]byte(config.Session.Secret))
	case consts.RedisStore:
		var err error
		sessionStore, err = redis.NewStoreWithDB(
			512,
			"tcp",
			config.Redis.Addrs[0],
			config.Redis.Password,
			strconv.Itoa(config.Session.DB),
			[]byte(config.Session.Secret),
		)
		if err != nil {
			return nil, err
		}
	default:
	}

	sessionStore.Options(sessions.Options{
		MaxAge:   config.Session.MaxAge,
		HttpOnly: config.Session.HttpOnly,
		Path:     "/",
	})

	return []gin.HandlerFunc{sessions.Sessions(SessionName, sessionStore)}, nil
}

func init() {
	AddMiddlewares(NewEntry(&Session{}, 0))
}
