package dto

import (
	"gorm.io/gorm"
)

type PagingDTO struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// NewPagingDTO 生成pagingDTO，默认为第一页，每页数据20条
func NewPagingDTO() PagingDTO {
	return PagingDTO{
		Page:     1,
		PageSize: 20,
	}
}

func (p *PagingDTO) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.PageSize
	}
	return offset
}

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
