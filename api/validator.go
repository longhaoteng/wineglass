package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	et "github.com/go-playground/validator/v10/translations/en"
	zt "github.com/go-playground/validator/v10/translations/zh"
	"golang.org/x/text/language"
)

var (
	enTrans ut.Translator
	zhTrans ut.Translator
)

type Bind struct {
	Fun func(obj interface{}) error
	Obj interface{}
}

func (a *API) Validator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(en.New(), zh.New())
		enTrans, _ = uni.GetTranslator("en")
		zhTrans, _ = uni.GetTranslator("zh")
		_ = et.RegisterDefaultTranslations(v, enTrans)
		_ = zt.RegisterDefaultTranslations(v, zhTrans)
	}
}

func (a *API) Verify(c *gin.Context, obj interface{}) (bool, *Response) {
	return a.verify(c, c.ShouldBind, obj)
}

func (a *API) Verifies(c *gin.Context, binds ...Bind) (bind bool, resp *Response) {
	for _, b := range binds {
		bind, resp = a.verify(c, b.Fun, b.Obj)
		if !bind {
			return bind, resp
		}
	}

	return
}

func GetLan(c *gin.Context) string {
	tags, _, _ := language.ParseAcceptLanguage(c.GetHeader("Accept-Language"))
	if len(tags) > 0 {
		str := strings.ToLower(tags[0].String())
		switch {
		case strings.Contains(str, "zh"):
			return "zh"
		default:
			return "en"
		}
	}
	return "en"
}

func (a *API) verify(c *gin.Context, fun func(obj interface{}) error, obj interface{}) (bool, *Response) {
	bind := false
	resp := &Response{}
	if err := fun(obj); err != nil {
		resp.HttpStatus = http.StatusBadRequest
		if vErrors, ok := err.(validator.ValidationErrors); ok {
			var trans ut.Translator
			switch GetLan(c) {
			case "zh":
				trans = zhTrans
			default:
				trans = enTrans
			}
			vErrs := vErrors.Translate(trans)
			errs := make(map[string]string)
			for s := range vErrs {
				errs[strings.ToLower(strings.Split(s, ".")[1])] = vErrs[s]
			}
			resp.Message = errs
		} else {
			resp.Message = err.Error()
		}
		a.Resp(c, resp)
	} else {
		bind = true
	}

	return bind, resp
}
