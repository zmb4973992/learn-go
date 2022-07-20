package model

import "time"

type OrganizationStructure struct {
	ID         int
	Name       string //组织结构名称
	level      string //组织结构级别，如公司、事业部、部门等
	CreatedAt  time.Time
	UpdatedAt  time.Time
	SuperiorID *int //上级机构ID，引用自身
	//这里是声名外键关系，并不是实际字段。结构体的字段名随意，首字母大写、否则不会导出，外键名会引用这个字段。
	//不建议用gorm的多对多的设定，不好修改
	Users       []User                  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SuperiorID1 []OrganizationStructure `gorm:"foreignkey:SuperiorID"` //设置外键规则，SuperiorID作为外键，引用自身ID
}

// TableName 修改表名
func (OrganizationStructure) TableName() string {
	return "organization_structure"
}
