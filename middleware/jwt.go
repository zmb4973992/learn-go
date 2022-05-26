package middleware

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/util/code"
	"learn-go/util/jwt"
	"net/http"
	"time"
)

func JWT() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authorization")
		if token == "" {
			c.JSON(http.StatusOK, serializer.CommonResponse{
				Data:    nil,
				Code:    code.Error_Token_Invalid,
				Message: code.GetErrorMessage(code.Error_Token_Invalid),
			})
			c.Abort()
			return
		}

		res, err := jwt.ParseToken(token)
		if err != nil || res.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusOK, serializer.CommonResponse{
				Data:    nil,
				Code:    code.Error_Token_Invalid,
				Message: code.GetErrorMessage(code.Error_Token_Invalid),
			})
			c.Abort()
			return
		}
		//这里只是用作测试jwt能否正常返回值，生产环境下只设置authorization、修改context、不返回任何信息，否则会对后续环节造成干扰
		//c.JSON(http.StatusOK, serializer.CommonResponse{
		//	Data: gin.H{
		//		"username": res.Username,
		//		"expire_at": time.Unix(res.ExpiresAt, 0).Format("2006-01-02 15:04:05"),
		//	},
		//	Code:    code.Success,
		//	Message: code.GetErrorMessage(code.Success),
		//})
		c.Set("username", res.Username)
		c.Next()
		return
	}
}
