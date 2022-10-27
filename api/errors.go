package api

import (
	"math"
	"net/http"
)

var (
	StatusTooManyRequests = &Error{
		Code: http.StatusTooManyRequests,
		Msg: ErrMsg{
			"en": http.StatusText(http.StatusTooManyRequests),
			"zh": "请求太频繁，请稍后再试",
		},
	}
)

// ErrMsg {"en": "user not found", "zh": "用户不存在"}
type ErrMsg map[string]string

type Error struct {
	Code  int
	Msg   ErrMsg
	Debug map[string]interface{}
}

type message struct {
	Msg   string                 `json:"msg"`
	Debug map[string]interface{} `json:"debug"`
}

func (e *Error) Error() string { return e.Msg[defaultLocale] }

func (e *Error) ErrCode() int { return e.Code }

func (e *Error) ErrMsg(locale string) interface{} {
	var msg string
	if m, ok := e.Msg[locale]; ok {
		msg = m
	} else {
		msg = e.Msg[defaultLocale]
	}

	if len(e.Debug) > 0 {
		return message{
			Msg:   msg,
			Debug: e.Debug,
		}
	}

	return msg
}

func (e *Error) DebugMsg(keysAndValues ...interface{}) {
	if len(keysAndValues) != 0 {
		e.Debug = make(map[string]interface{}, int(math.Ceil(float64(len(keysAndValues))/2)))
		for i := 0; i < len(keysAndValues); {
			key := keysAndValues[i]
			if keyStr, ok := key.(string); ok {
				if i+1 < len(keysAndValues) {
					e.Debug[keyStr] = keysAndValues[i+1]
				} else {
					e.Debug[keyStr] = ""
				}
			}
			i += 2
		}
	}
}

func ErrorWrap(err *Error, keysAndValues ...interface{}) error {
	err.DebugMsg(keysAndValues...)
	return err
}
