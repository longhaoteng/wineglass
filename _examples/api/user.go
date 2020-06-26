// @author mr.long

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/longhaoteng/wineglass"
	"github.com/longhaoteng/wineglass/_examples/err"
	"strings"
)

type User struct {
	api wineglass.API
}

type FetchUser struct {
	Name string
}

func (u *User) Router(r *gin.Engine) {
	r.GET("/user/:name", u.fetchUser)
}

func (u *User) fetchUser(c *gin.Context) {
	resp := &wineglass.Response{}
	if !strings.EqualFold(c.Param("name"), "wineglass") {
		u.api.Err(resp, err.UserNotFoundErr)
	}
	u.api.Resp(c, resp)
}
