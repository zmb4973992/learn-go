package dto

type UserDTO struct {
	ID       int      `json:"id"`
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Roles    []string `json:"roles"`
}

type UserGetDTO struct {
	Username          string   `json:"username"`            //用户名
	FullName          string   `json:"full_name"`           //全名
	EmailAddress      string   `json:"email_address"`       //邮箱地址
	IsValid           *bool    `json:"is_valid"`            //是否有效
	MobilePhoneNumber string   `json:"mobile_phone_number"` //手机号
	EmployeeNumber    string   `json:"employee_number"`     //工号
	Roles             []string `json:"roles"`               //角色
	Departments       []string `json:"departments"`         //部门
}

type UserListDTO struct {
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
