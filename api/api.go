package api

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/longhaoteng/wineglass/api/auth"
)

// Interface api interface
type Interface interface {
	Validator()
	API404(c *gin.Context)
	Verify(c *gin.Context, obj interface{}) (bool, *Response)
	Err(resp *Response, err error)
	Resp(c *gin.Context, r *Response)
}

type API struct{}

type Response struct {
	HttpStatus int
	Code       int
	Message    interface{}
	Data       interface{}
	Err        *Error
}

func (a *API) API404(c *gin.Context) {
	a.Resp(c, &Response{HttpStatus: http.StatusNotFound})
}

func (a *API) Err(resp *Response, err error) {
	switch e := err.(type) {
	case *Error:
		resp.Err = e
	default:
		resp.HttpStatus = http.StatusInternalServerError
	}
}

func (a *API) ErrResp(c *gin.Context, resp *Response, err error) {
	a.Err(resp, err)
	a.Resp(c, resp)
}

func (a *API) Resp(c *gin.Context, r *Response) {
	code := http.StatusOK
	if r == nil {
		r = &Response{}
	}
	if r.HttpStatus != 0 {
		code = r.HttpStatus
		r.Code = r.HttpStatus
	}
	if r.Err != nil {
		r.Code = r.Err.ErrCode()
		r.Message = r.Err.ErrMsg(GetLan(c))
	}
	if r.Code == 0 {
		r.Code = code
	}
	if r.Message == nil {
		r.Message = http.StatusText(code)
	}
	c.JSON(code, gin.H{
		"code":      r.Code,
		"msg":       r.Message,
		"data":      r.Data,
		"timestamp": time.Now().Unix(),
	})
}

func (a *API) Get(c *gin.Context, key interface{}) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

func (a *API) Set(c *gin.Context, key interface{}, val interface{}) error {
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

func (a *API) Delete(c *gin.Context, key interface{}) error {
	session := sessions.Default(c)
	session.Delete(key)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (a *API) SetToken(c *gin.Context, id int64, state bool) error {
	session := sessions.Default(c)
	session.Set(auth.TokenKey, &auth.User{ID: id, State: state})
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
