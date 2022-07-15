package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/status"
	"net/http"
)

func Login(c *gin.Context) {
	var s service.UserLoginService
	err := c.ShouldBind(&s)
	if err != nil {
		res := serializer.ResponseForDetail{
			Code:    status.ErrorInvalidJsonParameters,
			Data:    nil,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		}
		c.JSON(status.Error, res)
		return
	}
	res := s.Login()

	c.JSON(http.StatusOK, res)
}
