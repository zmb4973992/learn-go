package model

import "time"

type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Roles     []RoleAndUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName 修改表名
func (Role) TableName() string {
	return "role"
}
