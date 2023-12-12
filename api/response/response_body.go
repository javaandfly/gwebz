package response

import (
	"encoding/json"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Customize corresponding information
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    res.Data,
	}
}

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    data,
	}
}

func (res *Response) ToString() string {
	raw, _ := json.Marshal(&Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
	})
	return string(raw)
}

func response(code int, msg string) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}
