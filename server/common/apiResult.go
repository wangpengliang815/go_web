package common

import "encoding/json"

type ApiResponse struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// WithMsg 自定义响应信息
func (res *ApiResponse) WithMsg(message string) ApiResponse {
	return ApiResponse{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// WithData 追加响应数据
func (res *ApiResponse) WithData(data interface{}) ApiResponse {
	return ApiResponse{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
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

// 构造函数
func response(code int, msg string) *ApiResponse {
	return &ApiResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
