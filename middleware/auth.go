package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/status"
	"net/http"
)

func OnlyForAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, ok := c.Get("roles")
		if !ok {
			c.JSON(http.StatusForbidden, serializer.ResponseForDetail{
				Data:    nil,
				Code:    status.ErrorNeedAdminPrivilege,
				Message: status.GetMessage(status.ErrorNeedAdminPrivilege),
			})
			c.Abort()
			return
		}
		for _, role := range roles.([]string) {
			if role == "管理员权限" {
				c.Next()
				return
			}
			c.JSON(http.StatusForbidden, serializer.ResponseForDetail{
				Data:    nil,
				Code:    status.ErrorNeedAdminPrivilege,
				Message: status.GetMessage(status.ErrorNeedAdminPrivilege),
			})
			c.Abort()
			return
		}
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tempSubjects, ok := c.Get("roles") //获取用户角色,casbin规则的主体参数
		subjects := tempSubjects.([]string)
		if !ok {
			c.Abort()
			return
		}

		object := c.Request.URL.Path //获取请求路径，casbin规则的客体参数
		act := c.Request.Method      //获取请求方法，casbin规则的动作参数
		fmt.Println(subjects, object, act)
		e := util.NewEnforcer()
		if len(subjects) > 0 {
			for _, subject := range subjects {
				res, _ := e.Enforce(subject, object, act)
				if res {
					fmt.Println("通过")
					c.Next()
					break
				}
				fmt.Println("不通过")
				c.Abort()
				return
			}
		}
	}
}
