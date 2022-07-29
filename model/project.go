package model

import "time"

type Project struct {
	BaseModel
	ProjectCode      *string
	ProjectFullName  *string
	ProjectShortName *string
	Country          *string
	Province         *string
	ProjectType      *string
	Department       *string
	Amount           *int64
	Currency         *string
	ExchangeRate     *float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	RelatedPartyID   *int
	//外键，拆解情况
	Breakdowns []ProjectBreakdown `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName 将表名改为project
func (Project) TableName() string {
	return "project"
}
