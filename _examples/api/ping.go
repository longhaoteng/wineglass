package api

import (
	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass/api"
	"github.com/longhaoteng/wineglass/server"
)

type Ping struct {
	*api.API
}

func (p *Ping) Router(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		p.Data(c, "pong")
		p.Resp(c)
	})
}

func init() {
	server.AddRouters(&Ping{})
}
