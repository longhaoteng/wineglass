package errors

import (
	"github.com/longhaoteng/wineglass/api"
)

var (
	UserNotFoundErr = &api.Error{Code: 10000, Msg: api.ErrMsg{"en": "user not found", "zh": "用户不存在"}}
)
