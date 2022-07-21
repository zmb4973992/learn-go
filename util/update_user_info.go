package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpdateUserInfo(c *gin.Context, userID int) error {
	var roles []string
	c.Set("roles", []string{"vip1", "vip2"})

	temp, _ := c.Get("roles")
	roles = temp.([]string)
	for _, role := range roles {
		fmt.Println(role)
	}
	//报错：无法在多个赋值中将 any 赋给 roles (类型 []string)
}
