package common

type ApiResponse struct {
	Code    int         `json:"code"`    // 自定义错误码
	Message string      `json:"message"` // 错误描述
	Data    interface{} `json:"data"`    // 返回数据
}

// NewApiResponse 构造函数
func NewApiResponse(code int, message string, data interface{}) *ApiResponse {
	return &ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
