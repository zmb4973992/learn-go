package model

import "time"

type RoleAndUser struct {
	ID        int
	RoleID    int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 修改表名
func (RoleAndUser) TableName() string {
	return "role_and_user"
}
