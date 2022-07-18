package dto

type UserDTO struct {
	ID       int      `json:"id"`
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Roles    []string `json:"roles"`
}

type UserGetDTO struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
	Departments []string `json:"departments"` //待完善
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
