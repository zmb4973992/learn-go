package model

import "time"

type User struct {
	ID        uint64 `form:"id"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 将表名改为user
func (User) TableName() string {
	return "user"
}
