package model

import (
	"time"
)

type RelatedParty struct {
	ID                      int64
	ChineseName             *string
	EnglishName             *string
	SupplierCode            *string
	Address                 *string
	UniformSocialCreditCode *string //统一社会信用代码
	Telephone               *string
	File                    *string  `form:"file"`
	Files                   []string `form:"files"`
	CreatedAt               *time.Time
	UpdatedAt               *time.Time
	Project                 []Project `gorm:"foreignkey:RelatedPartyID;constraint:OnUpdate:CASCADE;"`
}

// TableName 将表名改为related_party
func (RelatedParty) TableName() string {
	return "related_party"
}
