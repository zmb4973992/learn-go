package model

import "time"

type User struct {
	ID        uint
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 将表名改为user
func (User) TableName() string {
	return "user"
}
