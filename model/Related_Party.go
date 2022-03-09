package model

import (
	"time"
)

type RelatedParty struct {
	ID                      uint64
	ChineseName             string
	EnglishName             string
	Relationship            string
	SupplierCode            string
	Address                 string
	UniformSocialCreditCode string //统一社会信用代码
	Telephone               string
	CreatedAt               time.Time
	UpdatedAt               time.Time
	Project                 []Project `gorm:"foreignkey:RelatedPartyID;constraint:OnUpdate:CASCADE;"`
}

func (RelatedParty) TableName() string {
	return "related_party"
}

type Project struct {
	ID               uint64
	RelatedPartyID   uint64
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
}

func (Project) TableName() string {
	return "project"
}
