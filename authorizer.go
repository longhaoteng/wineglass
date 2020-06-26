// @author mr.long

package wineglass

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	DefaultSessionKey = "role"
	// no session
	Anonymous = "anonymous"
	// expired session
	Unemployed = "unemployed"
)

func (m *Middleware) Authorizer() gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: m.Authorize.Enforcer, conf: m.Authorize}

	return func(c *gin.Context) {
		role := a.GetUserRole(c)
		if !a.CheckPermission(c, role) {
			switch role {
			case Anonymous:
				a.RequireLogIn(c)
			case Unemployed:
				a.RequireReLogIn(c)
			default:
				a.RequirePermission(c)
			}
		}
	}
}

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
	conf     *Authorize
}

func (a *BasicAuthorizer) GetUserRole(c *gin.Context) string {
	if _, err := c.Cookie("session"); err != nil {
		return Anonymous
	}
	session := sessions.Default(c)

	var sessionKey string
	if len(a.conf.SessionKey) > 0 {
		sessionKey = a.conf.SessionKey
	} else {
		sessionKey = DefaultSessionKey
	}

	role := session.Get(sessionKey)
	if role != nil {
		// rest ttl
		session.Set(sessionKey, role)
		if err := session.Save(); err != nil {
			log.Error(err)
		}
		return role.(string)
	}
	return Unemployed
}

func (a *BasicAuthorizer) CheckPermission(c *gin.Context, role string) bool {
	method := c.Request.Method
	path := c.Request.URL.Path

	allowed, err := a.enforcer.Enforce(role, path, method)
	if err != nil {
		panic(err)
	}

	return allowed
}

func (a *BasicAuthorizer) RequireLogIn(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}

func (a *BasicAuthorizer) RequireReLogIn(c *gin.Context) {
	c.AbortWithStatus(http.StatusProxyAuthRequired)
}

func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}
