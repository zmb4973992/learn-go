package dto

type DepartmentGetDTO struct {
	Name     string `json:"name"`     //部门名称
	Level    string `json:"level"`    //级别，如公司、事业部、部门等
	Superior any    `json:"superior"` //上级机构
}

type DepartmentUpdateDTO struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`        //部门名称
	Level      string `json:"level"`       //级别，如公司、事业部、部门等
	SuperiorID *int   `json:"superior_id"` //上级机构ID
}

//待实现
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
