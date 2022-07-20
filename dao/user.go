package dao

import (
	"fmt"
	"learn-go/dto"
	"learn-go/model"
	"learn-go/util"
)

type UserDAO struct{}

func (UserDAO) Get(id int) *dto.UserGetDTO {
	var tempUser model.User
	//先把基础的账号密码查出来
	err := model.DB.Model(&model.User{}).Where("id = ?", id).First(&tempUser).Error
	if err != nil {
		return nil
	}
	//然后把该userID的所有role_and_user记录查出来
	var roleAndUsers []model.RoleAndUser
	model.DB.Model(&model.RoleAndUser{}).Where("user_id = ?", id).Find(&roleAndUsers)
	//然后把所有的roleID提取出来，查出相应的角色名称
	var roles []string
	for _, record := range roleAndUsers {
		var role model.Role
		model.DB.Model(&model.Role{}).Where("id = ?", record.RoleID).Find(&role)
		roles = append(roles, role.Name)
	}
	//把所有查出的结果赋值给输出变量
	var u = dto.UserGetDTO{}
	u.ID = tempUser.ID
	u.Username = tempUser.Username
	u.Roles = roles
	var temp model.User
	model.DB.Debug().Preload("Department").Model(&model.User{}).Where("id = ?", 14).First(&temp)
	for _, v := range temp.Department {
		fmt.Println(*v.DepartmentID)
	}
	return &u
}

// Create 这里是只负责新增，不写任何业务逻辑。只要收到参数就创建数据库记录，然后返回错误
func (UserDAO) Create(paramIn *dto.UserDTO) error {
	user := new(model.User)
	user.Username = paramIn.Username
	user.Password = paramIn.Password
	err := model.DB.Create(user).Error
	return err
}

// Update 这里是只负责更新，不写任何业务逻辑。只要收到id和更新参数，然后返回错误
func (UserDAO) Update(id int, params map[string]any) error {
	//注意，这里就算没有找到记录，也不会报错，只有更新字段出现问题才会报错。详见gorm的update用法
	err := model.DB.Model(&model.User{}).Where("id = ?", id).Updates(params).Error
	return err
}

func (UserDAO) Delete(id int) error {
	//注意，这里就算没有找到记录，也不会报错。详见gorm的delete用法
	err := model.DB.Delete(&model.User{}, id).Error
	return err
}

// List 入参为sql查询条件，结果为数据列表+分页情况
func (UserDAO) List(sqlCondition util.SqlCondition) (
	list []dto.UserDTO, totalPages int, totalRecords int) {
	db := model.DB
	//select
	if len(sqlCondition.SelectedColumns) > 0 {
		db = db.Select(sqlCondition.SelectedColumns)
	}
	//where
	for _, paramPair := range sqlCondition.ParamPairs {
		db = db.Where(paramPair.ParamKey, paramPair.ParamValue)
	}
	//orderBy
	if sqlCondition.OrderBy.OrderByColumn != "" {
		if sqlCondition.OrderBy.Desc == true {
			db = db.Order(sqlCondition.OrderBy.OrderByColumn + " desc")
		} else {
			db = db.Order(sqlCondition.OrderBy.OrderByColumn)
		}
	}
	//count 计算totalRecords
	var tempTotalRecords int64
	err := db.Debug().Model(&model.User{}).Count(&tempTotalRecords).Error
	if err != nil {
		return nil, 0, 0
	}
	totalRecords = int(tempTotalRecords)

	//limit
	db = db.Limit(sqlCondition.Paging.PageSize)
	//offset
	offset := (sqlCondition.Paging.Page - 1) * sqlCondition.Paging.PageSize
	db = db.Offset(offset)

	//count 计算totalPages
	totalPages = model.GetTotalPages(totalRecords, sqlCondition.Paging.PageSize)
	err = db.Model(&model.User{}).Debug().Find(&list).Error
	if err != nil {
		return nil, 0, 0
	}
	return list, totalPages, totalRecords
}
