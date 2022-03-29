package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	libredis "github.com/go-redis/redis/v8"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"

	"github.com/longhaoteng/wineglass/api"
	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/consts"
	log "github.com/longhaoteng/wineglass/logger"
)

type Limiter struct{}

func (l *Limiter) Init() ([]gin.HandlerFunc, error) {
	var limiterStore limiter.Store

	storeOptions := limiter.StoreOptions{
		Prefix: "limiter",
	}

	switch config.Limiter.Store {
	case consts.MemoryStore:
		limiterStore = memory.NewStoreWithOptions(storeOptions)
	case consts.RedisStore:
		var err error
		client := libredis.NewClient(&libredis.Options{
			DB:           config.Redis.DB,
			Addr:         config.Redis.Addrs[0],
			PoolSize:     512,
			PoolTimeout:  10 * time.Second,
			IdleTimeout:  10 * time.Second,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			Password:     config.Redis.Password,
		})
		limiterStore, err = sredis.NewStoreWithOptions(client, storeOptions)
		if err != nil {
			return nil, err
		}
	default:
	}

	rate, err := limiter.NewRateFromFormatted(config.Limiter.Limit)
	if err != nil {
		return nil, err
	}

	return []gin.HandlerFunc{
		mgin.NewMiddleware(
			limiter.New(limiterStore, rate),
			mgin.WithErrorHandler(func(c *gin.Context, err error) {
				log.Field(consts.ErrKey, err).Log(log.ErrorLevel)
				c.Next()
			}),
			mgin.WithLimitReachedHandler(func(c *gin.Context) {
				code := http.StatusTooManyRequests
				c.JSON(code, gin.H{
					"code":      code,
					"data":      nil,
					"msg":       api.StatusTooManyRequests.ErrMsg(api.GetLanguage(c)),
					"timestamp": time.Now().Unix(),
				})
			}),
		),
	}, nil

}

func init() {
	AddMiddlewares(NewEntry(&Limiter{}))
}
