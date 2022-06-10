package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/util/status"
	"net/http"
)

type IBaseController interface {
	Success(c *gin.Context, data any)
	Failure(c *gin.Context, err status.CustomError)
}

type baseController struct {
}

func (b baseController) Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, serializer.ResponseForDetail{
		Data:    data,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	})
}

func (b baseController) Failure(c *gin.Context, err status.CustomError) {
	c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
		Data:    nil,
		Code:    err.ErrorCode,
		Message: err.ErrorMessage,
	})
}

func NewBaseController() IBaseController {
	return baseController{}
}
