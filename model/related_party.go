package model

import (
	"time"
)

type RelatedParty struct {
	ID                      int
	ChineseName             *string //中文名称
	EnglishName             *string //英文名称
	Address                 *string //地址
	UniformSocialCreditCode *string //统一社会信用代码
	Telephone               *string //电话
	File                    *string //单一文件的文件名，测试用，后期记得删
	Files                   *string //多个文件的文件名，测试用，后期记得删
	CreatedAt               *time.Time
	UpdatedAt               *time.Time
}

// TableName 修改数据库的表名
func (RelatedParty) TableName() string {
	return "related_party"
}
