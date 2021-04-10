package wineglass

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

const (
	// RateLimitMemoryStore Limiter Store use memory
	RateLimitMemoryStore = "memory"
	// RateLimitRedisStore Limiter Store use redis
	RateLimitRedisStore = "redis"
)

func (m *Middleware) RateLimiter() gin.HandlerFunc {
	conf := m.Limiter

	var rateLimiterStore limiter.Store

	storeOptions := limiter.StoreOptions{}
	if conf.Options != nil {
		storeOptions = *conf.Options

		if len(storeOptions.Prefix) == 0 {
			storeOptions.Prefix = "limiter"
		}
	}

	switch conf.Store {
	case RateLimitMemoryStore:
		rateLimiterStore = memory.NewStoreWithOptions(storeOptions)
	case RateLimitRedisStore:
		client := redis.NewClient(conf.Redis)

		var err error
		rateLimiterStore, err = sredis.NewStoreWithOptions(client, storeOptions)
		if err != nil {
			log.Fatalf("%+v\n", errors.New(err.Error()))
		}
	default:
		rateLimiterStore = memory.NewStoreWithOptions(storeOptions)
	}

	rate, err := limiter.NewRateFromFormatted(conf.Limit)
	if err != nil {
		log.Fatalf("%+v\n", errors.New(fmt.Sprintf("rate limiter conf.limit: %v", err.Error())))
	}

	return mgin.NewMiddleware(
		limiter.New(rateLimiterStore, rate),
		mgin.WithErrorHandler(func(c *gin.Context, err error) {
			fmt.Printf("%+v\n", errors.New(err.Error()))
			c.Next()
		}),
		mgin.WithLimitReachedHandler(func(c *gin.Context) {
			code := http.StatusTooManyRequests
			c.JSON(code, gin.H{
				"code":      code,
				"msg":       http.StatusText(code),
				"data":      nil,
				"timestamp": time.Now().Unix(),
			})
		}),
	)

}
