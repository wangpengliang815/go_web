package common

import "encoding/json"

type ApiResponse struct {
	Code int         `json:"code"` // 自定义错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// ToString 返回 JSON 格式的错误详情
func (res *ApiResponse) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// NewApiResponse 构造函数
func NewApiResponse(code int, msg string, data interface{}) *ApiResponse {
	return &ApiResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
