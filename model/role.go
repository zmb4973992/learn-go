package model

import "time"

type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	//这里是声名外键关系，并不是实际字段。不建议用gorm的多对多的设定，不好修改
	Users []RoleAndUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName 修改表名
func (Role) TableName() string {
	return "role"
}
