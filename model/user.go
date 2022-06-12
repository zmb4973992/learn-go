package model

import "time"

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Username  string
	Password  string
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
