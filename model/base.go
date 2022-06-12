package model

import "time"

type baseModel struct {
	ID        int        `json:"id" gorm`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
