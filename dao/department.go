package dao

import (
	"learn-go/dto"
	"learn-go/model"
	"learn-go/util"
)

func NewDepartmentDAO() DepartmentDAO {
	return DepartmentDAO{}
}

type DepartmentDAO struct{}

func (DepartmentDAO) Get(departmentID int) *dto.DepartmentGetDTO {
	var departmentGetDTO = dto.DepartmentGetDTO{}
	//把基础的部门信息查出来
	var department model.Department
	err := model.DB.Where("id = ?", departmentID).First(&department).Error
	if err != nil {
		return nil
	}
	//查询上级部门信息，采用递归方法
	if department.SuperiorID != nil {
		var tempDepartmentDAO DepartmentDAO
		departmentGetDTO.Superior = tempDepartmentDAO.Get(*department.SuperiorID)
	}
	//把所有查出的结果赋值给输出变量
	departmentGetDTO.Name = department.Name
	departmentGetDTO.Level = department.Level

	return &departmentGetDTO
}

// Create 这里是只负责新增，不写任何业务逻辑。只要收到参数就创建数据库记录，然后返回错误
func (DepartmentDAO) Create(paramIn *model.Department) error {
	err := model.DB.Create(paramIn).Error
	return err
}

// Update 这里是只负责更新，不写任何业务逻辑。只要收到id和更新参数，然后返回错误
func (DepartmentDAO) Update(id int, params map[string]any) error {
	//注意，这里就算没有找到记录，也不会报错，只有更新字段出现问题才会报错。详见gorm的update用法
	err := model.DB.Model(&model.User{}).Where("id = ?", id).Updates(params).Error
	return err
}

func (DepartmentDAO) Delete(id int) error {
	//注意，这里就算没有找到记录，也不会报错。详见gorm的delete用法
	err := model.DB.Delete(&model.User{}, id).Error
	return err
}

// List 入参为sql查询条件，结果为数据列表+分页情况
func (DepartmentDAO) List(sqlCondition util.SqlCondition) (
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
