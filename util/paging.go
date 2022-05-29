package util

import (
	"gorm.io/gorm"
)

type PagingRule struct {
	CurrentPage int `form:"current_page"`
	PageSize    int `form:"pagesize"`
}

func PaginateBy(rule PagingRule) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if rule.CurrentPage <= 0 {
			rule.CurrentPage = 1
		}
		switch {
		case rule.PageSize > 100:
			rule.PageSize = 1000
		case rule.PageSize <= 0:
			rule.PageSize = 20
		}
		offset := (rule.CurrentPage - 1) * rule.PageSize
		return db.Offset(offset).Limit(rule.PageSize)
	}
}

// GeneratePaginationRule 用于生成默认的分页器，默认返回第一页的数据，包含20条信息
func GeneratePaginationRule() *PagingRule {
	return &PagingRule{
		CurrentPage: 1,
		PageSize:    20,
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
