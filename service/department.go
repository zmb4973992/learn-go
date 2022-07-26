package service

import (
	"learn-go/dao"
	"learn-go/dto"
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

// Update 更新为什么要用dto？首先因为很多数据需要绑定，也就是一定要传参；
// 其次是需要清洗
func (DepartmentService) Update(paramIn *dto.DepartmentUpdateDTO) serializer.ResponseForDetail {
	var paramOut model.Department
	paramOut.ID = paramIn.ID
	paramOut.Name = paramIn.Name
	paramOut.Level = paramIn.Level
	//model.Department的SuperiorID为指针，需要处理
	if *paramIn.SuperiorID == 0 {
		paramOut.SuperiorID = nil
	} else {
		paramOut.SuperiorID = paramIn.SuperiorID
	}

	//清洗完毕，开始update
	r := dao.NewDepartmentDAO()
	err := r.Update(&paramOut)
	//拿到dao层的返回结果，进行处理
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
