package err

import "github.com/longhaoteng/wineglass"

var (
	UserNotFoundErr = &wineglass.Error{Code: 10000, Err: map[string]string{"en": "user not found", "zh": "用户不存在"}}
)
