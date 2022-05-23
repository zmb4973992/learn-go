package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/code"
	"net/http"
)

func Login(c *gin.Context) {
	var s service.UserLoginService
	err := c.ShouldBind(&s)
	if err != nil {
		res := serializer.CommonResponse{
			Code:    code.Error,
			Data:    "",
			Message: code.ErrorMessage[code.Error],
		}
		c.JSON(code.Error, res)
	}
	res := s.Login(s.Username)
	c.JSON(http.StatusOK, res)
}
