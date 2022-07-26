package dto

type DepartmentGetDTO struct {
	Name     string `json:"name"`     //部门名称
	Level    string `json:"level"`    //级别，如公司、事业部、部门等
	Superior any    `json:"superior"` //上级机构
}

// DepartmentUpdateDTO 只有update用，create直接用model
type DepartmentUpdateDTO struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`        //部门名称
	Level      string `json:"level" binding:"required"`       //级别，如公司、事业部、部门等
	SuperiorID *int   `json:"superior_id" binding:"required"` //上级机构ID
}

// DepartmentListDTO 待实现
type DepartmentListDTO struct {
	ID    int  `form:"id"`
	IDGte *int `form:"id_gte"`
	IDLte *int `form:"id_lte"`

	Username         string `form:"username"`
	UsernameNotEqual string `form:"username_ne"`
	UsernameInclude  string `form:"username_include"`
	Password         string `form:"password"`

	Paging  PagingDTO
	OrderBy OrderByDTO
}
