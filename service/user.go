package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/jwt"
	"learn-go/util/status"
)

type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserService struct {
	ID       int64   `form:"id" json:"id"`
	Username *string `form:"username" json:"username" binding:"required"`
	Password *string `form:"password" json:"-"  binding:"required"`
}

func (s *UserLoginService) Login() serializer.ResponseForDetail {
	var user model.User
	res := util.DB.Where("username=? and password=?", s.Username, s.Password).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return serializer.ResponseForDetail{
			Code:    status.ErrorUsernameOrPasswordExist,
			Message: status.GetMessage(status.ErrorUsernameOrPasswordExist),
		}
	}

	token := jwt.GenerateToken(*user.Username)
	return serializer.ResponseForDetail{
		Data: serializer.UserLoginResponse{
			Username: *user.Username,
			Token:    token,
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func GetUser(id int64) serializer.ResponseForDetail {
	var record *UserService
	result := util.DB.Debug().Model(&model.User{}).Where("id=?", id).First(&record)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	return serializer.ResponseForDetail{
		Data:    record,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func CreateUser(paramIn UserService) serializer.ResponseForDetail {
	record := new(model.User)
	if *paramIn.Username == "" || *paramIn.Password == "" {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorNotEnoughParameters,
			Message: status.GetMessage(status.ErrorNotEnoughParameters),
		}
	}
	record.Username = paramIn.Username
	record.Password = paramIn.Password
	res := util.DB.Debug().Create(&record)
	if res.Error != nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.Error,
			Message: status.GetMessage(status.Error),
		}
	}
	return serializer.ResponseForDetail{
		Data: UserService{
			ID:       record.ID,
			Username: record.Username,
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func UpdateUser(paramIn UserService) serializer.ResponseForDetail {
	record := &model.User{}
	res := util.DB.Debug().Model(&model.User{}).Where("id=?", paramIn.ID).First(&record)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	record.Username = paramIn.Username
	record.Password = paramIn.Password
	res = util.DB.Debug().Save(&record)
	if res.Error != nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorFailToSaveRecord,
			Message: status.GetMessage(status.ErrorFailToSaveRecord),
		}
	}
	return serializer.ResponseForDetail{
		Data:    nil,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func DeleteUser(id int64) serializer.ResponseForDetail {
	result := util.DB.Debug().Delete(&model.User{}, id)
	if result.Error != nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorFailToDeleteRecord,
			Message: status.GetMessage(status.ErrorFailToDeleteRecord),
		}
	}
	return serializer.ResponseForDetail{
		Data:    nil,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func GetUserList(paginationRule util.PagingRule) serializer.ResponseForDetail {
	var list []UserService
	fmt.Println(list)
	if paginationRule.CurrentPage <= 0 {
		paginationRule.CurrentPage = 1
	}
	if paginationRule.PageSize <= 0 || paginationRule.PageSize > 100 {
		paginationRule.PageSize = 20
	}
	util.DB.Debug().Scopes(util.PaginateBy(paginationRule)).Model(&model.User{}).Find(&list)
	return serializer.ResponseForDetail{
		Data:    list,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}
