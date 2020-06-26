// @author mr.long

package wineglass

type Error struct {
	Code int
	Err  string
}

func (e *Error) Error() string { return e.Err }

func (e *Error) ErrCode() int { return e.Code }
