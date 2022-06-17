package dto

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
