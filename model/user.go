package model

import "time"

type User struct {
	ID        int64  `form:"id" json:"id"`
	Username  string `form:"username" json:"username" `
	Password  string `form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 将表名改为user
func (User) TableName() string {
	return "user"
}
