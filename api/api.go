package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const ctxKey = "ctx"

type API struct {
	Session *Session
}

func NewApi() *API {
	return &API{
		Session: &Session{},
	}
}

func NoRoute(c *gin.Context) {
	a := NewApi()
	a.HttpStatus(c, http.StatusNotFound)
	a.Resp(c)
}

func (a *API) Context(c *gin.Context) *Context {
	apiCtx, exists := c.Get(ctxKey)
	if !exists {
		ctx := NewContext()
		c.Set(ctxKey, ctx)
		return ctx
	}
	return apiCtx.(*Context)
}

func (a *API) HttpStatus(c *gin.Context, status int) {
	a.Context(c).Response().SetHttpStatus(status)
}

func (a *API) Data(c *gin.Context, data interface{}) {
	a.Context(c).Response().SetData(data)
}

func (a *API) Verify(c *gin.Context, req interface{}) bool {
	return a.Verifies(c, req, c.ShouldBind)
}

func (a *API) Verifies(c *gin.Context, bindsAndFuncs ...interface{}) bool {
	if len(bindsAndFuncs) == 0 {
		return true
	}

	var binds []Bind
	for i := 0; i < len(bindsAndFuncs); {
		bindObj := bindsAndFuncs[i]
		if i+1 < len(bindsAndFuncs) {
			bindFunc := bindsAndFuncs[i+1]
			if f, ok := bindFunc.(func(obj interface{}) error); ok {
				binds = append(binds, Bind{
					Obj:  bindObj,
					Func: f,
				})
			}
		}
		i += 2
	}

	for _, bind := range binds {
		_, msg := bind.Verify(c)
		if msg != nil {
			ctx := a.Context(c)
			ctx.Response().SetHttpStatus(http.StatusBadRequest)
			ctx.Response().SetMsg(msg)
			a.Resp(c)
			return false
		}
	}
	return true
}

func (a *API) Groups(c *gin.Context, groups ...string) {
	ctx := a.Context(c)
	ctx.AddGroups(groups...)
}

func (a *API) Err(c *gin.Context, err error) {
	ctx := a.Context(c)
	switch e := err.(type) {
	case *Error:
		ctx.Response().SetErr(e)
	default:
		ctx.Response().SetHttpStatus(http.StatusInternalServerError)
	}
}

func (a *API) ErrResp(c *gin.Context, err error) {
	a.Err(c, err)
	a.Resp(c)
	c.Abort()
}

func (a *API) Resp(c *gin.Context) {
	if c.IsAborted() {
		return
	}

	ctx := a.Context(c)
	code := a.parseResp(c)
	if len(ctx.Groups()) == 0 {
		c.JSON(code, ctx.Response())
	} else {
		c.Render(code, DiffGroupsJSON{
			Groups: ctx.Groups(),
			Data:   ctx.Response(),
		})
	}
}

func (a *API) parseResp(c *gin.Context) int {
	code := http.StatusOK
	ctx := a.Context(c)
	resp := ctx.Response()
	resp.SetTime(time.Now().Unix())
	if resp.HttpStatus() != 0 {
		code = resp.HttpStatus()
		resp.Code = resp.HttpStatus()
	}
	if resp.Err() != nil {
		resp.SetCode(resp.Err().ErrCode())
		resp.SetMsg(resp.Err().ErrMsg(GetLanguage(c)))
	}
	if resp.GetCode() == 0 {
		resp.SetCode(code)
	}
	if resp.Msg() == nil {
		resp.SetMsg(http.StatusText(code))
	}

	return code
}
