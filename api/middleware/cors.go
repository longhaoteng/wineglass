package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/longhaoteng/wineglass/config"
)

type Cors struct{}

func (c *Cors) Init() ([]gin.HandlerFunc, error) {
	conf := cors.DefaultConfig()
	conf.AllowCredentials = true
	conf.AllowOrigins = config.Service.AllowOrigins
	return []gin.HandlerFunc{cors.New(conf)}, nil
}

func init() {
	AddMiddlewares(NewEntry(&Cors{}, 0))
}
