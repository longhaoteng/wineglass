package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass/api"
	"github.com/longhaoteng/wineglass/server"
)

type Hello struct {
	*api.API
}

type SayHello struct {
	Name string `binding:"required" json:"name"`
}

func (h *Hello) Router(r *gin.Engine) {
	r.GET("/hello/:name", h.getSay)
	r.POST("/hello", h.postSay)
}

func (h *Hello) getSay(c *gin.Context) {
	h.Data(c, fmt.Sprintf("hello %v", c.Param("name")))
	h.Resp(c)
}

func (h *Hello) postSay(c *gin.Context) {
	req := &SayHello{}
	// Verify required fields in the &SayHello{} struct, if name is empty, response 400.
	if h.Verify(c, req) {
		h.Data(c, fmt.Sprintf("hello %v", req.Name))
		h.Resp(c)
	}
	// multiple parameters
	// if h.Verifies(
	// 	c,
	// 	reqForPath, c.ShouldBindUri,
	// 	reqForBody, c.ShouldBind,
	// ) {}
}

func init() {
	server.AddRouters(&Hello{})
}
