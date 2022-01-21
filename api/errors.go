package api

import "net/http"

type Error struct {
	Code int
	// {"en": "user not found", "zh": "用户不存在"}
	Err map[string]string
}

func (e *Error) Error() string { return e.Err["en"] }

func (e *Error) ErrMsg(lan string) string { return e.Err[lan] }

func (e *Error) ErrCode() int { return e.Code }

var (
	StatusTooManyRequests = &Error{
		Code: http.StatusTooManyRequests,
		Err: map[string]string{
			"en": http.StatusText(http.StatusTooManyRequests),
			"zh": "请求太频繁，请稍后再试",
		},
	}
)
