package util

import (
	"gorm.io/gorm"
)

type Paging struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

func (p *Paging) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.PageSize
	}
	return offset
}

func PaginateBy(rule Paging) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if rule.Page <= 0 {
			rule.Page = 1
		}
		switch {
		case rule.PageSize > 100:
			rule.PageSize = 1000
		case rule.PageSize <= 0:
			rule.PageSize = 20
		}
		offset := (rule.Page - 1) * rule.PageSize
		return db.Offset(offset).Limit(rule.PageSize)
	}
}

// GeneratePaginationRule 用于生成默认的分页器，默认返回第一页的数据，包含20条信息
func GeneratePaginationRule() *Paging {
	return &Paging{
		Page:     1,
		PageSize: 20,
	}
}

// GetTotalPage 使用前请确保两个参数均大于0，否则结果会返回0
func GetTotalPage(totalNumberOfRecord int, pageSize int) (totalPage int) {
	if totalNumberOfRecord <= 0 || pageSize <= 0 {
		return 0
	}
	totalPage = totalNumberOfRecord / pageSize
	if totalNumberOfRecord%pageSize != 0 {
		totalPage++
	}
	return
}
