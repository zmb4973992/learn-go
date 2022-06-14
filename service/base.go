package service

import (
	"learn-go/serializer"
	"learn-go/util/status"
)

type baseService struct{}

func (baseService) Success(data any) any {
	return serializer.ResponseForDetail{
		Data:    data,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func (baseService) Failure(errCode int) any {
	return serializer.ResponseForDetail{
		Data:    nil,
		Code:    errCode,
		Message: status.GetMessage(errCode),
	}
}
