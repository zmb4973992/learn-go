package service

import (
	"learn-go/dao"
	"learn-go/dto"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/status"
	"strconv"
	"time"
)

/*
Service层没有数据结构、只有方法，所有的数据结构都放在DTO里
入参为id、DTO，出参为response。
这里的方法从controller拿来id或初步处理的入参dto，重点是处理业务逻辑。
所有的增删改查都交给DAO层处理，否则service层会非常庞大。
生成出参response后，交给controller展示。
*/

type RelatedPartyService struct {
	baseService
	ID                      int
	ChineseName             *string       `form:"chinese_name"`
	EnglishName             *string       `form:"english_name" `
	SupplierCode            *string       `form:"supplier_code" `
	Address                 *string       `form:"address" `
	UniformSocialCreditCode *string       `form:"uniform_social_credit_code" ` //统一社会信用代码
	Telephone               *string       `form:"telephone" `
	File                    *string       `form:"-"`
	Files                   *string       `form:"-"`
	CreatedAt               *time.Time    `form:"created_at"`
	UpdatedAt               *time.Time    `form:"updated_at"`
	Paging                  dto.PagingDTO `json:"-"`
}

func NewRelatedPartyService() RelatedPartyService {
	return RelatedPartyService{}
}

func (RelatedPartyService) Get(id int) *serializer.ResponseForDetail {
	u := new(dao.RelatedPartyDAO)
	result := u.Get(id)
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

func (RelatedPartyService) List(paramIn dto.RelatedPartyListDTO) serializer.ResponseForList {
	//生成sql查询条件
	sqlCondition := util.NewSqlCondition()
	//对paramIn进行清洗
	//这部分是用于where的参数
	if paramIn.Paging.Page > 0 {
		sqlCondition.Paging.Page = paramIn.Paging.Page
	}
	//如果参数里的pageSize是整数且大于0、小于等于100：
	if paramIn.Paging.PageSize > 0 && paramIn.Paging.PageSize <= 100 {
		sqlCondition.Paging.PageSize = paramIn.Paging.PageSize
	}
	if id := paramIn.ID; id > 0 {
		sqlCondition.Equal("id", id)
	}
	if idGte, err := strconv.Atoi(paramIn.IDGte); err == nil {
		sqlCondition.Gte("id", idGte)
	}
	if idLte, err := strconv.Atoi(paramIn.IDLte); err == nil {
		sqlCondition.Lte("id", idLte)
	}
	if paramIn.ChineseName != "" {
		sqlCondition = sqlCondition.Equal("chinese_name", paramIn.ChineseName)
	}
	if paramIn.ChineseNameInclude != "" {
		sqlCondition = sqlCondition.Include("chinese_name", paramIn.ChineseNameInclude)
	}

	//这部分是用于order的参数
	column := paramIn.OrderBy.OrderByColumn
	//allColumns := []string{"id", "telephone", "file"}
	//re := util.IsInSlice(column, allColumns)
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
	u := new(dao.RelatedPartyDAO)
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
			Page:         sqlCondition.Paging.Page,
			PageSize:     sqlCondition.Paging.PageSize,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func (RelatedPartyService) Update(paramIn *model.RelatedParty) serializer.ResponseForDetail {
	//var record model.RelatedParty
	//对model进行清洗，生成dao层需要的model

	//清洗完毕，开始update
	r := dao.NewRelatedDAO()
	err := r.Update(paramIn)
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

func (RelatedPartyService) Create(paramIn *model.RelatedParty) serializer.ResponseForDetail {
	//对model进行清洗，生成dao层需要的model

	//清洗完毕，开始create
	r := dao.NewRelatedDAO()
	err := r.Create(paramIn)
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

func (RelatedPartyService) Delete(id int) serializer.ResponseForDetail {
	r := dao.NewRelatedDAO()
	err := r.Delete(id)
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
