package model

import (
	"time"
)

type Test struct {
	ID        int64
	Str       *string    `form:"str"`
	Num       *int       `form:"num"`
	Time      *time.Time `form:"time"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// TableName 将表名改为related_party
func (Test) TableName() string {
	return "test"
}
