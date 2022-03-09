package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// MyMiddlewareTest 这里用闭包的形式来写中间件，也是gin推荐的方式
func MyMiddlewareTest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("我是中间件1")
	}
}
