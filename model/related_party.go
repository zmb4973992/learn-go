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
	File                    *string //单一文件的文件名
	Files                   *string //多个文件的文件名
	CreatedAt               *time.Time
	UpdatedAt               *time.Time
	Project                 []Project `gorm:"foreignkey:RelatedPartyID;constraint:OnUpdate:CASCADE;"`
}

// TableName 将表名改为related_party
func (RelatedParty) TableName() string {
	return "related_party"
}
