package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"learn-go/dao"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/jwt"
	"learn-go/util/status"
)

type UserLoginService struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s UserLoginService) Login() serializer.ResponseForDetail {
	var user model.User
	//根据入参的用户名，从数据库取出记录赋值给user
	res := dao.DB.Where("username=?", s.Username).First(&user)
	//如果没有找到记录
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		response := serializer.NewErrorResponse(status.ErrorInvalidUsernameOrPassword)
		return response
	}
	//如果找到记录了，但是密码错误的话
	if util.CheckPassword(s.Password, user.Password) == false {
		response := serializer.NewErrorResponse(status.ErrorInvalidUsernameOrPassword)
		return response
	}
	//账号密码都正确时，生成token
	token := jwt.GenerateToken(user.ID)
	return serializer.ResponseForDetail{
		Data: gin.H{
			"access_token": token,
			"roles":        []string{"公司级权限", "事业部级权限"},
			"说明":           "这是写死在service/login.go的内容",
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}
