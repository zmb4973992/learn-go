package dao

import "learn-go/model"

type DepartmentAndUserDAO struct{}

// Create 这里是只负责新增，不写任何业务逻辑。只要收到参数就创建数据库记录，然后返回错误
func (DepartmentAndUserDAO) Create(param *model.DepartmentAndUser) error {
	err := model.DB.Create(param).Error
	return err
}

func (DepartmentAndUserDAO) CreateBatch(param []model.DepartmentAndUser) error {
	err := model.DB.Create(param).Error
	return err
}
