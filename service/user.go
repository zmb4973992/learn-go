package service

import (
	"errors"
	"gorm.io/gorm"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
)

type UserLoginService struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (s *UserLoginService) Login(username string) serializer.CommonResponse {
	var user model.User
	res := util.DB.Where("username=?", s.Username).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return serializer.CommonResponse{Message: "错误"}
	}

	util.DB.Debug().Exec("delete from `user` where username = ?", "333")

	return serializer.CommonResponse{Data: "66"}

}
