package service

import (
	"errors"
	"gorm.io/gorm"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/code"
	"learn-go/util/jwt"
)

type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (s *UserLoginService) Login() serializer.CommonResponse {
	var user model.User
	res := util.DB.Where("username=? and password=?", s.Username, s.Password).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return serializer.CommonResponse{
			Code:    code.Error_Username_Or_Password_Exist,
			Message: code.GetErrorMessage(code.Error_Username_Or_Password_Exist),
		}
	}

	token := jwt.GenerateToken(user.Username)
	return serializer.CommonResponse{
		Data: serializer.UserLoginResponse{
			Username: user.Username,
			Token:    token,
		},
		Code:    code.Success,
		Message: code.GetErrorMessage(code.Success),
	}
}
