package service

import (
	"learn-go/dao"
	"learn-go/dto"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util/status"
)

// ProjectBreakdownService 没有数据、只有方法，所有的数据都放在DTO里
//这里的方法从controller拿来初步处理的入参，重点是处理业务逻辑
//所有的增删改查都交给DAO层处理，否则service层会非常庞大
type ProjectBreakdownService struct {
	baseService
}

func (ProjectBreakdownService) Get(projectBreakdownID int) *serializer.ResponseForDetail {
	u := new(dao.ProjectBreakdownDAO)
	result := u.Get(projectBreakdownID)
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

func (ProjectBreakdownService) Create(paramIn *dto.ProjectBreakdownCreateAndUpdateDTO) serializer.ResponseForDetail {
	//对dto进行清洗，生成dao层需要的model
	var paramOut model.ProjectBreakdown
	paramOut.Name = paramIn.Name
	paramOut.Level = paramIn.Level
	paramOut.ProjectID = paramIn.ProjectID

	//model.Department的SuperiorID为指针，需要处理
	if *paramIn.SuperiorID == 0 {
		paramOut.SuperiorID = nil
	} else {
		paramOut.SuperiorID = paramIn.SuperiorID
	}
	d := new(dao.DepartmentDAO)
	err := d.Create(&paramOut)
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
