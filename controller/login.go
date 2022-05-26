package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/code"
	"learn-go/util/jwt"
	"net/http"
)

func Login(c *gin.Context) {
	var s service.UserLoginService
	err := c.ShouldBind(&s)
	if err != nil {
		res := serializer.CommonResponse{
			Code:    code.Error,
			Data:    nil,
			Message: code.GetErrorMessage(code.Error),
		}
		c.JSON(code.Error, res)
		return
	}

	res := s.Login()
	token := jwt.GenerateToken(s.Username)
	c.Header("authorization", token)
	c.JSON(http.StatusOK, res)
}

func ParseTokenTest(c *gin.Context) {
	u, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"n": u,
	})

}
