package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/util/status"
	"net/http"
)

type IBaseController interface {
	Success(c *gin.Context, data any)
	Failure(c *gin.Context, errCode int)
}

type baseController struct{}

func (baseController) Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, serializer.ResponseForDetail{
		Data:    data,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	})
}

func (baseController) Failure(c *gin.Context, errCode int) {
	response := serializer.NewErrorResponse(errCode)
	c.JSON(http.StatusBadRequest, response)
}
