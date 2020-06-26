// @author mr.long

package err

import "github.com/longhaoteng/wineglass"

var (
	UserNotFoundErr = &wineglass.Error{Code: 10000, Err: "user not found"}
)
