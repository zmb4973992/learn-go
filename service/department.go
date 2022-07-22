package service

import (
	"learn-go/dao"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util/status"
)

// DepartmentService 没有数据、只有方法，所有的数据都放在DTO里
//这里的方法从controller拿来初步处理的入参，重点是处理业务逻辑
//所有的增删改查都交给DAO层处理，否则service层会非常庞大
type DepartmentService struct {
	baseService
}

func NewDepartmentService() DepartmentService {
	return DepartmentService{}
}

func (s DepartmentService) Get(departmentID int) *serializer.ResponseForDetail {
	u := new(dao.DepartmentDAO)
	result := u.Get(departmentID)
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

func (DepartmentService) Create(paramIn *model.Department) serializer.ResponseForDetail {
	//对model进行清洗，生成dao层需要的model

	d := dao.NewDepartmentDAO()
	err := d.Create(paramIn)
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
