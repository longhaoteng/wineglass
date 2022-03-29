package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass/_examples/errors"
	"github.com/longhaoteng/wineglass/api"
	"github.com/longhaoteng/wineglass/server"
)

type User struct {
	*api.API
}

func (u *User) Router(r *gin.Engine) {
	r.GET("/user/:name", u.fetchUser)
}

func (u *User) fetchUser(c *gin.Context) {
	if !strings.EqualFold(c.Param("name"), "wineglass") {
		u.ErrResp(c, errors.UserNotFoundErr)
	}
	u.Resp(c)
}

func init() {
	server.AddRouters(&User{})
}
