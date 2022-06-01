package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/jwt"
	"learn-go/util/status"
	"net/http"
)

func Login(c *gin.Context) {
	var s service.UserService
	err := c.ShouldBind(&s)
	if err != nil {
		res := serializer.ResponseForDetail{
			Code:    status.ErrorInvalidFormDataParameters,
			Data:    nil,
			Message: status.GetMessage(status.ErrorInvalidFormDataParameters),
		}
		c.JSON(status.Error, res)
		return
	}

	res := s.Login()
	token := jwt.GenerateToken(*s.Username)
	c.Header("authorization", token)
	c.JSON(http.StatusOK, res)
}
