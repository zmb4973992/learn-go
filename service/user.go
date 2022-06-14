package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"learn-go/dao"
	"learn-go/dto"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/status"
)

// UserService 没有数据、只有方法，所有的数据都放在DTO里
//这里的方法从controller拿来初步处理的入参，重点是处理业务逻辑
//所有的增删改查都交给DAO层处理，否则service层会非常庞大
type UserService struct {
	baseService
}

func NewUserService() UserService {
	return UserService{}
}

func (s UserService) Get(id int) (data any) {
	u := new(dao.UserDAO)
	data = u.Get(id)
	return data
}

func (UserService) Create(paramIn *dto.UserDTO) serializer.ResponseForDetail {
	encryptedPassword, err := util.EncryptPassword(paramIn.Password)
	if err != nil {
		return serializer.NewErrorResponse(status.ErrorFailToEncrypt)
	}
	paramIn.Password = encryptedPassword

	var userDAO dao.UserDAO
	err = userDAO.Create(paramIn)
	if err != nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.Error,
			Message: status.GetMessage(status.Error),
		}
	}
	return serializer.ResponseForDetail{
		Data:    nil,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func (UserService) Update(id int, paramIn map[string]any) serializer.ResponseForDetail {
	//新建一个dao.User结构体的实例
	u := new(dao.UserDAO)
	//然后使用它的方法
	err := u.Update(id, paramIn)
	//如果返回的错误是“没找到指定记录”：
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	//如果是其他错误：
	if err != nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorFailToSaveRecord,
			Message: status.GetMessage(status.ErrorFailToSaveRecord),
		}
	}
	//更新正常、没有错误的话：
	return serializer.ResponseForDetail{
		Data:    nil,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func (UserService) Delete(id int) serializer.ResponseForDetail {
	//新建一个dao.User结构体的实例
	u := new(dao.UserDAO)
	err := u.Delete(id)
	if err != nil {
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

func (UserService) List(paramIn dto.UserListDTO) serializer.ResponseForDetail {
	//生成sql查询条件
	sqlCondition := util.NewSqlCondition()
	//对paramIn进行清洗
	//这部分是用于where的参数
	page := paramIn.Paging.Page
	if page > 0 {
		sqlCondition.Paging.Page = page
	}
	//如果参数里的pageSize是整数且大于0、小于等于100：
	pageSize := paramIn.Paging.PageSize
	if pageSize > 0 && pageSize <= 100 {
		sqlCondition.Paging.PageSize = pageSize
	}
	id := paramIn.ID
	if id > 0 {
		sqlCondition.Where("id", id)
	}

	idGte := paramIn.IDGte
	if idGte != nil && *idGte >= 0 {
		sqlCondition.Gte("id", *idGte)
	}
	idLte := paramIn.IDLte
	if idLte != nil && *idLte >= 0 {
		sqlCondition.Lte("id", *idLte)
	}
	username := paramIn.Username
	if username != "" {
		sqlCondition = sqlCondition.Where("username = ?", username)
	}
	password := paramIn.Password
	if password != "" {
		sqlCondition = sqlCondition.Where("password = ?", password)
	}
	usernameNoteEqual := paramIn.UsernameNotEqual
	if usernameNoteEqual != "" {
		sqlCondition = sqlCondition.NotEqual("username", usernameNoteEqual)
	}
	usernameInclude := paramIn.UsernameInclude
	if usernameInclude != "" {
		sqlCondition = sqlCondition.Include("username", usernameInclude)
	}

	//这部分是用于order的参数
	column := paramIn.OrderBy.OrderByColumn
	if column != "" {
		sqlCondition.OrderBy.OrderByColumn = column
	}
	desc := paramIn.OrderBy.Desc
	if desc == true {
		sqlCondition.OrderBy.Desc = true
	} else {
		sqlCondition.OrderBy.Desc = false
	}
	fmt.Println(sqlCondition.Paging.Page)
	fmt.Println(sqlCondition.Paging.PageSize)
	//新建一个dao.User结构体的实例
	u := new(dao.UserDAO)
	list := u.List(*sqlCondition)
	return serializer.ResponseForDetail{
		Data:    list,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}
