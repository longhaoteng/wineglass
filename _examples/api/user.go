package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass/_examples/errors"
	"github.com/longhaoteng/wineglass/api"
	"github.com/longhaoteng/wineglass/server"
)

type User struct {
	api api.API
}

func (u *User) Router(r *gin.Engine) {
	r.GET("/user/:name", u.fetchUser)
}

func (u *User) fetchUser(c *gin.Context) {
	resp := &api.Response{}
	if !strings.EqualFold(c.Param("name"), "wineglass") {
		u.api.Err(resp, errors.UserNotFoundErr)
	}
	u.api.Resp(c, resp)
}

func init() {
	server.AddRouters(&User{})
}
