package model

import (
	"time"
)

type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	//这里是声名外键关系，并不是实际字段。不建议用gorm的多对多的设定，不好修改
	User []RoleAndUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName 修改表名
func (Role) TableName() string {
	return "role"
}

func generateRoles() error {
	roles := []Role{
		{Name: "管理员"},
		{Name: "公司级"},
		{Name: "事业部级"},
		{Name: "部门级"},
		{Name: "项目级"},
	}
	for _, role := range roles {
		err := DB.FirstOrCreate(&Role{}, role).Error
		if err != nil {
			return err
		}
	}
	return nil
}
