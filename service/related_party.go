package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"learn-go/dao"
	"learn-go/dto"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util/status"
	"time"
)

type RelatedPartyService struct {
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

func GetDetailOfRelatedParty(id int64) (*serializer.ResponseForDetail, error) {
	var record *model.RelatedParty
	result := dao.DB.Debug().Where("id=?", id).First(&record)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}, nil
	}
	return &serializer.ResponseForDetail{
		Data:    record,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}, nil
}

func UpdateRelatedParty(paramIn RelatedPartyService) serializer.ResponseForDetail {
	var record model.RelatedParty
	result := dao.DB.Debug().First(&record, paramIn.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	record.ChineseName = paramIn.ChineseName
	record.EnglishName = paramIn.EnglishName
	record.Address = paramIn.Address
	record.UniformSocialCreditCode = paramIn.UniformSocialCreditCode
	record.Telephone = paramIn.Telephone
	result = dao.DB.Debug().Save(&record)
	if result.Error != nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.Error,
			Message: status.GetMessage(status.Error),
		}
	}
	return serializer.ResponseForDetail{
		Data:    nil,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func CreateRelatedParty(paramIn RelatedPartyService) serializer.ResponseForDetail {
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
	if paramIn.File != nil && *paramIn.File != "" {
		record.File = paramIn.File
	}
	if paramIn.Files != nil && *paramIn.Files != "" {
		record.Files = paramIn.Files
	}
	result := dao.DB.Debug().Create(&record)
	if result.Error != nil {
		return serializer.NewResponseForCreationResult(status.ErrorFailToSaveRecord, record.ID)
	}
	return serializer.NewResponseForCreationResult(status.Success, record.ID)
}

func DeleteRelatedParty(id int64) serializer.ResponseForDetail {
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
