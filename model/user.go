package model

import "time"

type User struct {
	ID                int `json:"id"`
	Username          string
	Password          string
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
	IsValid           bool      //用户为有效还是禁用
	FullName          *string
	EmailAddress      *string
	MobilePhoneNumber *string
	JobNumber         *string
	//这里是声名外键关系，并不是实际字段。不建议用gorm的多对多的设定，不好修改
	//角色
	Roles       []RoleAndUser       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Departments []DepartmentAndUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName 将表名改为user
func (User) TableName() string {
	return "user"
}

//func UserExistOrNot(username string) (code uint64) {
//	var user model.User
//	util.DB.Where("username = ?", username).First(&user)
//	if user.ID > 0 {
//		return util.Error_Username_Exist
//	} else {
//		return user.ID
//	}
//}
