package serializer

import "learn-go/util/status"

type ResponseForDetail struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseForList struct {
	Data    any    `json:"data"`
	Paging  any    `json:"paging"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseForPaging struct {
	CurrentPage int `json:"page"`
	PageSize    int `json:"page_size"`
	TotalPage   int `json:"total_page"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// NewResponseForCreationResult 创建记录后，自动生成返回结果
func NewResponseForCreationResult(statusCode int, ID int) ResponseForDetail {
	return ResponseForDetail{
		Data: map[string]any{
			"ID": ID,
		},
		Code:    statusCode,
		Message: status.GetMessage(statusCode),
	}
}

// NewErrorResponse 创建记录后，自动生成返回结果
func NewErrorResponse(statusCode int) ResponseForDetail {
	return ResponseForDetail{
		Data:    nil,
		Code:    statusCode,
		Message: status.GetMessage(statusCode),
	}
}
