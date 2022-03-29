package api

import (
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

const defaultLocale = "en"

var (
	trans       = make(map[string]ut.Translator)
	localeAlias = make(map[string]string)
)

type BindFunc func(obj interface{}) error

type Bind struct {
	Obj  interface{}
	Func BindFunc
}

func AddTranslator(locale string, t *ut.Translator, alias ...string) {
	trans[locale] = *t

	localeAlias[locale] = locale
	for _, a := range alias {
		localeAlias[strings.ToLower(a)] = locale
	}
}

func AddLocaleAlias(locale, alias string) {
	localeAlias[strings.ToLower(alias)] = locale
}

func GetLanguage(c *gin.Context) string {
	tags, _, _ := language.ParseAcceptLanguage(c.GetHeader("Accept-Language"))
	if len(tags) > 0 {
		tag := strings.ToLower(tags[0].String())
		if locale, ok := localeAlias[tag]; ok {
			return locale
		}
	}
	return defaultLocale
}

func Validator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(en.New(), zh.New())
		enTrans, _ := uni.GetTranslator("en")
		zhTrans, _ := uni.GetTranslator("zh")
		if err := et.RegisterDefaultTranslations(v, enTrans); err != nil {
			return err
		}
		if err := zt.RegisterDefaultTranslations(v, zhTrans); err != nil {
			return err
		}
		AddTranslator("en", &enTrans, "en-US", "en-GB")
		AddTranslator("zh", &zhTrans, "zh-CN")
	}
	return nil
}

func (b *Bind) Verify(c *gin.Context) (bool, interface{}) {
	if err := b.Func(b.Obj); err != nil {
		if vErrors, ok := err.(validator.ValidationErrors); ok {
			translator := trans[defaultLocale]
			if locale, ok := localeAlias[GetLanguage(c)]; ok {
				if t, ok := trans[locale]; ok {
					translator = t
				}
			}
			vErrs := vErrors.Translate(translator)
			errs := make(map[string]string)
			for s := range vErrs {
				errs[strings.ToLower(strings.Split(s, ".")[1])] = vErrs[s]
			}
			return false, errs
		} else {
			return false, err.Error()
		}
	}

	return true, nil
}
