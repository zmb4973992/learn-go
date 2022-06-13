package dao

import (
	"fmt"
	"learn-go/dto"
	"learn-go/model"
	"learn-go/util"
)

type UserDAO struct{}

func (UserDAO) Get(id int) *dto.UserDTO {
	var u *dto.UserDTO
	err := DB.Model(&model.User{}).Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil
	}
	return u
}

// Create 这里是只负责新增，不写任何业务逻辑。只要收到参数就创建数据库记录，然后返回错误
func (UserDAO) Create(paramIn *dto.UserDTO) error {
	user := new(model.User)
	user.Username = paramIn.Username
	user.Password = paramIn.Password
	err := DB.Create(user).Error
	return err
}

// Update 这里是只负责更新，不写任何业务逻辑。只要收到id和更新参数，然后返回错误
func (UserDAO) Update(id int, params map[string]any) error {
	//注意，这里就算没有找到记录，也不会报错，只有更新字段出现问题才会报错。详见gorm的update用法
	err := DB.Model(&model.User{}).Where("id = ?", id).Updates(params).Error
	return err
}

func (UserDAO) Delete(id int) error {
	//注意，这里就算没有找到记录，也不会报错。详见gorm的delete用法
	err := DB.Delete(&model.User{}, id).Error
	return err
}

// List 入参为sql查询条件，结果为数据列表+分页情况
func (UserDAO) List(sqlCondition util.SqlCondition) (list []dto.UserDTO) {
	db := DB
	//select
	if len(sqlCondition.SelectedColumns) > 0 {
		db = db.Select(sqlCondition.SelectedColumns)
	}
	//where
	for _, paramPair := range sqlCondition.ParamPairs {
		db = db.Where(paramPair.ParamKey, paramPair.ParamValue)
	}
	//orderBy
	var order string
	if sqlCondition.OrderBy.Ascending == true {
		order = ""
	} else {
		order = " desc"
	}
	column := sqlCondition.OrderBy.Column
	if column != "" {
		db = db.Order(column + order)
	}
	//limit

	//offset

	//db.Model(&model.User{}).Debug().Find(&list)
	err := sqlCondition.Find(db, &model.User{}, &list)
	if err != nil {
		fmt.Println(err)
	}
	return list
}
