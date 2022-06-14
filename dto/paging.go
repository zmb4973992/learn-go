package dto

import (
	"gorm.io/gorm"
)

type PagingDTO struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type OrderByDTO struct {
	OrderByColumn string `form:"order_by"` //排序字段
	Desc          bool   `form:"desc"`     //是否为降序（从大到小）
}

//func (p *PagingDTO) Offset() int {
//	offset := 0
//	if p.Page > 0 {
//		offset = (p.Page - 1) * p.PageSize
//	}
//	return offset
//}

// PaginateBy 拟废除，不要再用了
func PaginateBy(rule PagingDTO) func(db *gorm.DB) *gorm.DB {
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

// GetTotalPage 使用前请确保两个参数均大于0，否则结果会返回0。拟废除，不要再用了
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
