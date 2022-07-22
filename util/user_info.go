package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/model"
)

func UpdateUserInfo(c *gin.Context, userID int) {
	var user model.User
	model.DB.Preload("Roles").First(&user)
	var roleNames []string
	for _, role := range user.Roles {
		var tempRole model.Role
		model.DB.Where("id = ?", role.RoleID).First(&tempRole)
		roleNames = append(roleNames, tempRole.Name)
	}
	fmt.Println(roleNames)
}
