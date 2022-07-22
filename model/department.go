package model

type Department struct {
	BaseModel
	Name       string `json:"name" binding:"required"`  //名称
	Level      string `json:"level" binding:"required"` //层级，如公司、事业部、部门等
	SuperiorID *int   `json:"superior_id"`              //上级机构ID，引用自身
	//这里是声名外键关系，并不是实际字段。结构体的字段名随意，首字母大写、否则不会导出，外键名会引用这个字段。
	//不建议用gorm的多对多的设定，不好修改

	//设置外键规则，SuperiorID作为外键，引用自身ID
	//数据库规则限制，自引用不能设置级联更新和级联删除
	SuperiorID1 []Department        `gorm:"foreignkey:SuperiorID"`
	Users       []DepartmentAndUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName 修改表名
func (Department) TableName() string {
	return "department"
}
