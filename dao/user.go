package dao

import (
	"learn-go/model"
	"learn-go/util"
)

func CreateUser(user *model.User) (code uint64) {
	code = UserExistOrNot(user.Username)
	if code == util.Success {
		err := util.DB.Create(&user).Error
		if err != nil {
			return util.Error
		} else {
			return util.Success
		}
	} else {
		return util.Error
	}
}

func UserExistOrNot(username string) (code uint64) {
	var user model.User
	util.DB.Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return util.Error_Username_Exist
	} else {
		return user.ID
	}
}
