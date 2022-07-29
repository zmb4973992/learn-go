package service

import (
	"errors"
	"gorm.io/gorm"
	"learn-go/dao"
	"learn-go/dto"
	"learn-go/model"
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

func (UserService) Get(userID int) *serializer.ResponseForDetail {
	u := new(dao.UserDAO)
	result := u.Get(userID)
	if result == nil {
		return &serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	return &serializer.ResponseForDetail{
		Data:    result,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func (UserService) Create(paramIn *dto.UserCreateDTO) serializer.ResponseForDetail {
	//把基础数据添加到user表
	var paramOutForUser model.User
	paramOutForUser.Username = paramIn.Username
	//对密码进行加密
	encryptedPassword, err := util.EncryptPassword(paramIn.Password)
	if err != nil {
		return serializer.NewErrorResponse(status.ErrorFailToEncrypt)
	}
	paramOutForUser.Password = encryptedPassword
	paramOutForUser.IsValid = paramIn.IsValid
	if *paramIn.FullName == "" {
		paramOutForUser.FullName = nil
	} else {
		paramOutForUser.FullName = paramIn.FullName
	}
	if *paramIn.EmailAddress == "" {
		paramOutForUser.EmailAddress = nil
	} else {
		paramOutForUser.EmailAddress = paramIn.EmailAddress
	}
	if *paramIn.MobilePhoneNumber == "" {
		paramOutForUser.MobilePhoneNumber = nil
	} else {
		paramOutForUser.MobilePhoneNumber = paramIn.MobilePhoneNumber
	}
	if *paramIn.EmployeeNumber == "" {
		paramOutForUser.EmployeeNumber = nil
	} else {
		paramOutForUser.EmployeeNumber = paramIn.EmployeeNumber
	}

	//由于涉及到多表的保存，所以这里启用事务
	err = model.DB.Transaction(func(tx *gorm.DB) error {
		//注意，这里没有使用dao层的封装方法，而是使用tx+gorm的原始方法
		err = tx.Create(&paramOutForUser).Error
		if err != nil {
			return err
		}
		//把用户-角色的对应关系添加到role_and_user表
		//如果有角色数据：
		if len(paramIn.Roles) > 0 {
			var paramOutForRoleAndUser []model.RoleAndUser

			//这里不能使用v进行循环赋值，因为涉及到指针，会导致所有记录都变成一样的
			for k := range paramIn.Roles {
				var record model.RoleAndUser
				record.UserID = &paramOutForUser.ID
				record.RoleID = &paramIn.Roles[k]
				paramOutForRoleAndUser = append(paramOutForRoleAndUser, record)
			}
			err = tx.Create(paramOutForRoleAndUser).Error
			if err != nil {
				return err
			}
		}

		//把用户-部门的对应关系添加到department_and_user表
		//如果有部门数据：
		if len(paramIn.Departments) > 0 {
			var paramOutForDepartmentAndUser []model.DepartmentAndUser
			for k := range paramIn.Departments {
				var record model.DepartmentAndUser
				record.UserID = &paramOutForUser.ID
				record.DepartmentID = &paramIn.Departments[k]
				paramOutForDepartmentAndUser = append(paramOutForDepartmentAndUser, record)
			}
			err = tx.Create(paramOutForDepartmentAndUser).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
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

func (UserService) List(paramIn dto.UserListDTO) serializer.ResponseForList {
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
	//新建一个dao.User结构体的实例
	u := new(dao.UserDAO)
	list, totalPages, totalRecords := u.List(*sqlCondition)
	if list == nil {
		return serializer.ResponseForList{
			Data:    nil,
			Paging:  nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	return serializer.ResponseForList{
		Data: list,
		Paging: &dto.PagingDTO{
			Page:         page,
			PageSize:     pageSize,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}
