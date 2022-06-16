package dto

import (
	"gorm.io/gorm"
)

// PagingDTO 处理分页相关的数据
//只接收form(query)形式的page和pageSize
//发送时通过json全部发送
type PagingDTO struct {
	Page         int `form:"page"      json:"page"`
	PageSize     int `form:"page_size" json:"page_size"`
	TotalPages   int `form:"-"         json:"total_pages"`
	TotalRecords int `form:"-"         json:"total_records"`
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
