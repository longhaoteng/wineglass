package api

type Response struct {
	httpStatus int
	err        *Error
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Message    interface{} `json:"message"`
	Timestamp  int64       `json:"timestamp"`
}

func (r *Response) SetHttpStatus(status int) {
	r.httpStatus = status
}

func (r *Response) HttpStatus() int {
	return r.httpStatus
}

func (r *Response) SetErr(err *Error) {
	r.err = err
}

func (r *Response) Err() *Error {
	return r.err
}

func (r *Response) GetCode() int {
	return r.Code
}

func (r *Response) SetCode(code int) {
	r.Code = code
}

func (r *Response) GetData() interface{} {
	return r.Data
}

func (r *Response) SetData(data interface{}) {
	r.Data = data
}

func (r *Response) Msg() interface{} {
	return r.Message
}

func (r *Response) SetMsg(msg interface{}) {
	r.Message = msg
}

func (r *Response) Time() int64 {
	return r.Timestamp
}

func (r *Response) SetTime(timestamp int64) {
	r.Timestamp = timestamp
}
