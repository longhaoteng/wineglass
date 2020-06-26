// @author mr.long

package wineglass

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"net/http"
	"time"
)

const (
	// Limiter Store use memory
	RateLimitMemoryStore = "memory"
	// Limiter Store use redis
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
		if storeOptions.MaxRetry == 0 {
			storeOptions.MaxRetry = 3
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
			log.Fatal(err)
		}
	default:
		rateLimiterStore = memory.NewStoreWithOptions(storeOptions)
	}

	rate, err := limiter.NewRateFromFormatted(conf.Limit)
	if err != nil {
		log.Fatalf("rate limiter conf.limit: %v",err)
	}

	return mgin.NewMiddleware(
		limiter.New(rateLimiterStore, rate),
		mgin.WithErrorHandler(func(c *gin.Context, err error) {
			log.Errorln(err)
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
