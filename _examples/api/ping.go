package api

import (
	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass/api"
	"github.com/longhaoteng/wineglass/server"
)

type Ping struct {
	api api.API
}

func (p *Ping) Router(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		p.api.Resp(c, &api.Response{Data: "pong"})
	})
}

func init() {
	server.AddRouters(&Ping{})
}
