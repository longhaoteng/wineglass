// @author mr.long

package wineglass

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/ulule/limiter/v3"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

// Wineglass config
type Config struct {
	// RunMode sets gin mode according to input string.
	RunMode string
	// Pprof serves
	Pprof *Pprof
	// Gin Middleware
	Middleware *Middleware
}

type Pprof struct {
	// Open pprof serves
	Open bool
	// Default path prefix is "/debug/pprof"
	Prefix string
}

type Middleware struct {
	// Cors middleware, when it is empty not use.
	Cors *cors.Config
	// Session middleware, when it is empty not use.
	Session *Session
	// Authorizer middleware, when it is empty not use, prerequisites for session middleware activation.
	Authorize *Authorize
	// RateLimiter middleware, when it is empty not use.
	Limiter *RateLimit
	// Other middlewares
	Middlewares []gin.HandlerFunc
}

type Session struct {
	// Whether to enable this middleware.
	Enable bool
	// Session store
	Store string
	// Default "session"
	Name string
	// Session secret string.
	Secret string
	// Options stores configuration for a session or session store.
	// Fields are a subset of http.Cookie fields.
	Options *sessions.Options
	// Register global module
	GobModels []interface{}
	// Session store use redis.
	// Ref: https://godoc.org/github.com/boj/redistore#NewRediStoreWithDB
	Redis *struct {
		Size                  int
		Address, Password, DB string
	}
}

type Authorize struct {
	// Whether to enable this middleware.
	Enable bool
	// Enforcer is the main interface for authorization enforcement and policy management.
	Enforcer *casbin.Enforcer
	// Default "role"
	SessionKey string
}

type RateLimit struct {
	// Whether to enable this middleware.
	Enable bool
	// Limiter store
	Store string
	// Limiter are options for store.
	Options *limiter.StoreOptions
	// format:<limit>-<period>
	// 5 reqs/second: "5-S"
	// 10 reqs/minute: "10-M"
	// 1000 reqs/hour: "1000-H"
	// 2000 reqs/day: "2000-D"
	Limit string
	// RateLimiter use redis.
	Redis *redis.Options
}
