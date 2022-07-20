package model

import "time"

type Project struct {
	ID               int64
	ProjectCode      string
	ProjectFullName  string
	ProjectShortName string
	Country          string
	Province         string
	ProjectType      string
	Department       string
	Amount           int64
	Currency         string
	ExchangeRate     float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	//外键
	RelatedPartyID int
}

// TableName 将表名改为project
func (Project) TableName() string {
	return "project"
}
