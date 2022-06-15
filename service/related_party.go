package service

import (
	"fmt"
	"learn-go/dao"
	"learn-go/dto"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util/status"
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

func GetRelatedPartyList(s RelatedPartyService) serializer.ResponseForList {
	if s.Paging.Page <= 0 {
		s.Paging.Page = 1
	}
	if s.Paging.PageSize <= 0 || s.Paging.PageSize > 100 {
		s.Paging.PageSize = 20
	}
	var list []RelatedPartyService
	res := dao.DB.Debug().Scopes(dto.PaginateBy(s.Paging)).
		Model(&model.RelatedParty{}).
		Where("chinese_name=?", s.ChineseName).
		Find(&list)
	var temp int64
	dao.DB.Debug().
		Model(&model.RelatedParty{}).
		Where("chinese_name=?", s.ChineseName).Count(&temp)
	if res.Error != nil {
		return serializer.ResponseForList{
			Data:    nil,
			Paging:  nil,
			Code:    status.Error,
			Message: status.GetMessage(status.Error),
		}
	}
	if res.RowsAffected == 0 {
		return serializer.ResponseForList{
			Data:    nil,
			Paging:  nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	fmt.Println(res.RowsAffected)
	fmt.Println(temp)
	return serializer.ResponseForList{
		Data:    list,
		Paging:  nil,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func (RelatedPartyService) Update(paramIn *dto.RelatedPartyDTO) serializer.ResponseForDetail {
	var record model.RelatedParty
	//对dto进行清洗，生成dao层需要的model
	if &paramIn.ChineseName != nil {
		record.ChineseName = paramIn.ChineseName
	}
	if paramIn.EnglishName != nil {
		record.EnglishName = paramIn.EnglishName
	}
	if paramIn.Address != nil {
		record.Address = paramIn.Address
	}
	if paramIn.UniformSocialCreditCode != nil {
		record.UniformSocialCreditCode = paramIn.UniformSocialCreditCode
	}
	if paramIn.Telephone != nil {
		record.Telephone = paramIn.Telephone
	}
	record.ID = paramIn.ID
	//清洗完毕，开始update
	r := dao.NewRelatedDAO()
	err := r.Update(&record)
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

func (RelatedPartyService) Create(paramIn *dto.RelatedPartyDTO) serializer.ResponseForDetail {
	//对dto进行清洗，生成dao层需要的model
	var record model.RelatedParty
	if paramIn.ChineseName != nil && *paramIn.ChineseName != "" {
		record.ChineseName = paramIn.ChineseName
	}
	if paramIn.EnglishName != nil && *paramIn.EnglishName != "" {
		record.EnglishName = paramIn.EnglishName
	}
	if paramIn.Address != nil && *paramIn.Address != "" {
		record.Address = paramIn.Address
	}
	if paramIn.UniformSocialCreditCode != nil && *paramIn.UniformSocialCreditCode != "" {
		record.UniformSocialCreditCode = paramIn.UniformSocialCreditCode
	}
	if paramIn.Telephone != nil && *paramIn.Telephone != "" {
		record.Telephone = paramIn.Telephone
	}
	//清洗完毕，开始create
	r := dao.NewRelatedDAO()
	err := r.Create(&record)
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
	result := dao.DB.Debug().Delete(&model.RelatedParty{}, id)
	if result.Error != nil {
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
