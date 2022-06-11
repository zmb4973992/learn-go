package middleware

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/util/jwt"
	"learn-go/util/status"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("access_token")
		//如果请求头没有携带access_token
		if token == "" {
			c.JSON(http.StatusOK, serializer.ResponseForDetail{
				Data:    nil,
				Code:    status.ErrorAccessTokenNotFound,
				Message: status.GetMessage(status.ErrorAccessTokenNotFound),
			})
			c.Abort()
			return
		}
		//开始校验access_token
		res, err := jwt.ParseToken(token)
		if err != nil || res.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusOK, serializer.ResponseForDetail{
				Data:    nil,
				Code:    status.ErrorAccessTokenInvalid,
				Message: status.GetMessage(status.ErrorAccessTokenInvalid),
			})
			c.Abort()
			return
		}
		//如果access_token校验通过
		c.Set("username", res.Username)
		c.Next()
		return
		//这里只是用作测试jwt能否正常返回值，生产环境下只设置authorization、修改context、不返回任何信息，否则会对后续环节造成干扰
		//c.JSON(http.StatusOK, serializer.ResponseForDetail{
		//	Data: gin.H{
		//		"username": res.Username,
		//		"expire_at": time.Unix(res.ExpiresAt, 0).Format("2006-01-02 15:04:05"),
		//	},
		//	Code:    code.Success,
		//	Message: code.GetMessage(code.Success),
		//})

	}
}
