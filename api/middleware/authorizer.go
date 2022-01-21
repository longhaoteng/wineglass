package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/longhaoteng/wineglass/api/auth"
	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/conv"
	log "github.com/longhaoteng/wineglass/logger"
)

const (
	// Masked session invalid
	Masked = "masked"
	// Guest no session
	Guest = "guest"
)

type Authorizer struct{}

func (a *Authorizer) Init() ([]gin.HandlerFunc, error) {
	b := &BasicAuthorizer{}
	return []gin.HandlerFunc{
		func(c *gin.Context) {
			id, _, role := b.GetUser(c)
			allow, err := b.CheckPermission(c.Request, id)
			if err != nil {
				log.Error(err)
				b.TryAgainLater(c)
				return
			}

			if !allow {
				switch role {
				case Guest:
					b.RequireLogIn(c)
				case Masked:
					b.RequireReLogIn(c)
				default:
					b.RequirePermission(c)
				}
			}
		},
	}, nil
}

type BasicAuthorizer struct{}

func (a *BasicAuthorizer) GetUser(c *gin.Context) (id int64, state bool, role string) {
	if _, err := c.Cookie(SessionName); err != nil {
		return 0, true, Guest
	}

	session := sessions.Default(c)
	user := session.Get(auth.TokenKey)
	if user != nil {
		// rest ttl
		session.Set(SessionName, user)
		if err := session.Save(); err != nil {
			log.Error(err)
		}

		u := user.(*auth.User)
		return u.ID, u.State, ""
	}

	return 0, true, Masked
}

func (a *BasicAuthorizer) CheckPermission(r *http.Request, id int64) (bool, error) {
	allow, err := auth.Enforce(
		conv.FormatInt64(id),
		config.Service.Name,
		r.URL.Path,
		r.Method,
	)
	if err != nil {
		return false, err
	}

	return allow, nil
}

func (a *BasicAuthorizer) TryAgainLater(c *gin.Context) {
	a.Abort(c, http.StatusBadGateway, "try again later")
}

func (a *BasicAuthorizer) RequireLogIn(c *gin.Context) {
	a.Abort(c, http.StatusUnauthorized, "need to login")
}

func (a *BasicAuthorizer) RequireReLogIn(c *gin.Context) {
	a.Abort(c, 499, "invalid login")
}

func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	a.Abort(c, http.StatusForbidden, "no permission")
}

func (a *BasicAuthorizer) Abort(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":      code,
		"msg":       msg,
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})
}

func init() {
	AddMiddlewares(NewEntry(&Authorizer{}, 1))
}
