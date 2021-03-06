package dao

import (
	"learn-go/dto"
	"learn-go/model"
	"learn-go/util"
)

type UserDAO struct{}

func (UserDAO) Get(userID int) *dto.UserGetDTO {
	//把基础的账号信息查出来
	var user model.User
	err := model.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil
	}
	//把该userID的所有role_and_user记录查出来
	var roleAndUsers []model.RoleAndUser
	model.DB.Where("user_id = ?", userID).Find(&roleAndUsers)
	//把所有的roleID提取出来，查出相应的角色名称
	var roleNames []string
	for _, roleAndUser := range roleAndUsers {
		var role model.Role
		model.DB.Where("id = ?", roleAndUser.RoleID).First(&role)
		roleNames = append(roleNames, role.Name)
	}
	//把该userID的所有department_and_user记录查出来
	var departmentAndUsers []model.DepartmentAndUser
	model.DB.Where("user_id = ?", userID).Find(&departmentAndUsers)
	//把所有的departmentID提取出来，查出相应的部门名称
	var departmentNames []string
	for _, departmentAndUser := range departmentAndUsers {
		var department model.Department
		model.DB.Where("id = ?", departmentAndUser.DepartmentID).First(&department)
		departmentNames = append(departmentNames, department.Name)
	}
	//把所有查出的结果赋值给输出变量
	var userGetDTO = dto.UserGetDTO{}
	userGetDTO.Username = user.Username
	if user.IsValid != nil {
		userGetDTO.IsValid = user.IsValid
	}
	if user.FullName != nil {
		userGetDTO.FullName = *user.FullName
	}
	if user.EmailAddress != nil {
		userGetDTO.EmailAddress = *user.EmailAddress
	}
	if user.MobilePhoneNumber != nil {
		userGetDTO.MobilePhoneNumber = *user.MobilePhoneNumber
	}
	if user.EmployeeNumber != nil {
		userGetDTO.EmployeeNumber = *user.EmployeeNumber
	}

	userGetDTO.Roles = roleNames
	userGetDTO.Departments = departmentNames
	return &userGetDTO
}

// Create 这里是只负责新增，不写任何业务逻辑。只要收到参数就创建数据库记录，然后返回错误
func (UserDAO) Create(param *model.User) error {
	err := model.DB.Create(param).Error
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
	list []dto.UserCreateDTO, totalPages int, totalRecords int) {
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
	err := db.Model(&model.User{}).Count(&tempTotalRecords).Error
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
	err = db.Model(&model.User{}).Find(&list).Error
	if err != nil {
		return nil, 0, 0
	}
	return list, totalPages, totalRecords
}
