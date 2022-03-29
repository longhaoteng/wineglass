package api

import "net/http"

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
	Code int
	Msg  ErrMsg
}

func (e *Error) Error() string { return e.Msg[defaultLocale] }

func (e *Error) ErrMsg(locale string) string {
	if msg, ok := e.Msg[locale]; ok {
		return msg
	}

	return e.Msg[defaultLocale]
}

func (e *Error) ErrCode() int { return e.Code }
