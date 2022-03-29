package api

import (
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"

	"github.com/longhaoteng/wineglass/api/auth"
	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/conv"
)

type Session struct{}

func (s *Session) Get(c *gin.Context, key interface{}) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

func (s *Session) Set(c *gin.Context, key interface{}, val interface{}) error {
	if reflect.ValueOf(val).IsNil() {
		return nil
	}
	session := sessions.Default(c)
	session.Set(key, val)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (s *Session) Delete(c *gin.Context, key interface{}) error {
	session := sessions.Default(c)
	session.Delete(key)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (s *Session) SetToken(c *gin.Context, id int64, state bool) error {
	return s.Set(c, auth.TokenKey, &auth.User{ID: id, State: state})
}

func (s *Session) GetToken(c *gin.Context) *auth.User {
	user := s.Get(c, auth.TokenKey)
	if u, ok := user.(*auth.User); ok {
		return u
	}
	return &auth.User{}
}

func (s *Session) GetRoles(c *gin.Context) []string {
	return auth.Enforcer().GetRolesForUserInDomain(
		conv.FormatInt64(s.GetToken(c).ID),
		config.Service.Name,
	)
}

func (s *Session) HasRole(c *gin.Context, role string) bool {
	return funk.ContainsString(s.GetRoles(c), role)
}
