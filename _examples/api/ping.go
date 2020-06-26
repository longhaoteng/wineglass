// @author mr.long

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass"
)

type Ping struct {
	api wineglass.API
}

func (p *Ping) Router(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		p.api.Resp(c, &wineglass.Response{Data: "pong"})
	})
}
