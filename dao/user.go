package dao

import "learn-go/model"

type user struct {
}

func Create(user *model.User) error {
	err := DB.Create(&user).Error
	return err
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
