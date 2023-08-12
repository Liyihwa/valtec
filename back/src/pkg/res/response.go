package res

/*
Response 是前后端沟通的媒介
code目前仅有200(成功),201(失败)两种
msg为原因
data为存放数据的键值对
*/
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data_   any    `json:"data,omitempty"`
}

func Ok() *Response {
	return &Response{Code: 200, Message: "", Data_: nil}
}
func Fail() *Response {
	return &Response{Code: 201, Message: "", Data_: nil}
}

func (r *Response) Data(value any) *Response {
	r.Data_ = value
	return r
}
func (r *Response) Msg(msg string) *Response {
	r.Message = msg
	return r
}
