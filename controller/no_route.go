package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/util/status"
	"net/http"
)

type NoRouteController struct {
}

func NewNoRouteController() NoRouteController {
	return NoRouteController{}
}

func (NoRouteController) NoRoute(c *gin.Context) {
	c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
		Data:    nil,
		Code:    status.ErrorInvalidRequest,
		Message: status.GetMessage(status.ErrorInvalidRequest),
	})
}
