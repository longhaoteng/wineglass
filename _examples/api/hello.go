package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass/api"
	"github.com/longhaoteng/wineglass/server"
)

type Hello struct {
	api api.API
}

type SayHello struct {
	Name string `binding:"required" json:"name"`
}

func (h *Hello) Router(r *gin.Engine) {
	r.GET("/hello/:name", h.getSay)
	r.POST("/hello", h.postSay)
}

func (h *Hello) getSay(c *gin.Context) {
	h.api.Resp(c, &api.Response{Data: fmt.Sprintf("hello %v", c.Param("name"))})
}

func (h *Hello) postSay(c *gin.Context) {
	req := &SayHello{}
	// Verify required fields in the &SayHello{} struct, if name is empty, response 400.
	if bind, resp := h.api.Verify(c, req); bind {
		resp.Data = fmt.Sprintf("hello %v", req.Name)
		h.api.Resp(c, resp)
	}
}

func init() {
	server.AddRouters(&Hello{})
}
