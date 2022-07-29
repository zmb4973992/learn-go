package util

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"learn-go/model"
)

func SetUserInfo(c *gin.Context, userID int) {
	var user model.User
	//预加载全部关联信息
	model.DB.Where("id = ?", userID).Preload(clause.Associations).First(&user)
	//设置拥有的权限
	var roleNames []string
	for _, role := range user.Roles {
		var roleInfo model.Role
		model.DB.Where("id = ?", role.RoleID).First(&roleInfo)
		roleNames = append(roleNames, roleInfo.Name)
	}
	c.Set("roles", roleNames)
	//设置所属部门
	var departmentNames []string
	for _, department := range user.Departments {
		var departmentInfo model.Department
		model.DB.Where("id = ?", department.DepartmentID).First(&departmentInfo)
		departmentNames = append(departmentNames, departmentInfo.Name)
	}
	c.Set("departments", departmentNames)
}
