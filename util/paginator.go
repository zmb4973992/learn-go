package util

import (
	"gorm.io/gorm"
)

type PaginationRule struct {
	Page     int `form:"page"`
	PageSize int `form:"pagesize"`
}

func Paginate(p *PaginationRule) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page <= 0 {
			p.Page = 1
		}
		switch {
		case p.PageSize > 100:
			p.PageSize = 1000
		case p.PageSize <= 0:
			p.PageSize = 20
		}
		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

// GeneratePaginationRule 用于生成默认的分页器，默认返回第一页的数据，包含20条信息
func GeneratePaginationRule() *PaginationRule {
	return &PaginationRule{
		Page:     1,
		PageSize: 20,
	}
}

func GetTotalPage(totalRecord int, pageSize int) (totalPage int) {
	totalPage = totalRecord / pageSize
	if totalRecord%pageSize != 0 {
		totalPage++
	}
	return
}
