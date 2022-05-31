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
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"pagesize"`
	TotalPage   int `json:"total_page"`
}

type UserLoginResponse struct {
	Username string
	Token    string
}

func NewResponseForDetail(statusCode int) ResponseForDetail {
	return ResponseForDetail{
		Data:    nil,
		Code:    statusCode,
		Message: status.GetMessage(statusCode),
	}
}
