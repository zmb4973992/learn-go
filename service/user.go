package service

import (
	"errors"
	"gorm.io/gorm"
	"learn-go/dao"
	"learn-go/dto"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/status"
)

type UserService struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserService() UserService {
	return UserService{}
}

func (UserService) Get(id int) any {
	u := new(dao.UserDAO)
	data := u.Get(id)
	if data == nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	return serializer.ResponseForDetail{
		Data:    data,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
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

func (UserService) List(paramIn map[string]any) serializer.ResponseForDetail {
	//生成sql查询条件
	sqlCondition := util.NewSqlCondition()
	//对paramIn进行清洗
	//这部分是用于where的参数
	//如果类型为any，那么go会把数字识别为float64，需要在这里进行转化
	page, ok := paramIn["page"].(float64)
	if ok && page > 0 {
		sqlCondition.Paging.Page = int(page)
	}
	//如果参数里的pageSize是整数且大于0、小于等于100：
	pageSize, ok := paramIn["page_size"].(float64)
	if ok && pageSize > 0 && pageSize <= 100 {
		sqlCondition.Paging.PageSize = int(pageSize)
	}
	if paramIn["username"] != "" {
		sqlCondition = sqlCondition.Where("username = ?", paramIn["username"])
	}
	if paramIn["password"] != "" {
		sqlCondition = sqlCondition.Where("password = ?", paramIn["password"])
	}
	id, ok := paramIn["id_gte"].(float64)
	if ok && id > 0 {
		sqlCondition = sqlCondition.Where("id >= ?", int(id))
	}
	//这部分是用于order的参数
	orderBy, ok := paramIn["order_by"].(string)
	if ok {
		ascending, ok := paramIn["ascending"].(bool)
		//要考虑前端只传orderBy字段、没传顺序方向字段
		if ok { //如果传了顺序方向，就用传来的值
			sqlCondition.OrderBy.Column = orderBy
			sqlCondition.OrderBy.Ascending = ascending
		} else { //如果没传顺序方向或读取错误，就默认为升序，从小到大
			sqlCondition.OrderBy.Column = orderBy
			sqlCondition.OrderBy.Ascending = true
		}
	}

	//新建一个dao.User结构体的实例
	u := new(dao.UserDAO)
	list := u.List(*sqlCondition)
	return serializer.ResponseForDetail{
		Data:    list,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}
