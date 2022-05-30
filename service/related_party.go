package service

import (
	"errors"
	"gorm.io/gorm"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/status"
	"mime/multipart"
	"time"
)

type RelatedPartyService struct {
	ID                      int64
	ChineseName             *string               `form:"chinese_name"`
	EnglishName             *string               `form:"english_name" `
	SupplierCode            *string               `form:"supplier_code" `
	Address                 *string               `form:"address" `
	UniformSocialCreditCode *string               `form:"uniform_social_credit_code" ` //统一社会信用代码
	Telephone               *string               `form:"telephone" `
	File                    *multipart.FileHeader `form:"file123"`
	CreatedAt               *time.Time            `form:"created_at"`
	UpdatedAt               *time.Time            `form:"updated_at"`
}

func GetRelatedPartyList(paginationRule util.PagingRule) serializer.ResponseForList {
	var list []RelatedPartyService
	if paginationRule.CurrentPage <= 0 {
		paginationRule.CurrentPage = 1
	}
	if paginationRule.PageSize <= 0 || paginationRule.PageSize > 100 {
		paginationRule.PageSize = 20
	}
	res := util.DB.Debug().Scopes(util.PaginateBy(paginationRule)).Model(&model.RelatedParty{}).Find(&list)
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
	var total int64
	util.DB.Debug().Model(&model.RelatedParty{}).Find(&list).Count(&total)
	return serializer.ResponseForList{
		Data: list,
		Paging: serializer.ResponseForPaging{
			CurrentPage: paginationRule.CurrentPage,
			PageSize:    paginationRule.PageSize,
			TotalPage:   util.GetTotalPage(int(total), paginationRule.PageSize),
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func GetDetailOfRelatedParty(id int64) (*serializer.ResponseForDetail, error) {
	var record *model.RelatedParty
	result := util.DB.Debug().Where("id=?", id).First(&record)
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

func UpdateDetailOfRelatedParty(paramIn RelatedPartyService) serializer.ResponseForDetail {
	var record model.RelatedParty
	result := util.DB.Debug().First(&record, paramIn.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		}
	}
	record.ChineseName = paramIn.ChineseName
	record.EnglishName = paramIn.EnglishName
	record.SupplierCode = paramIn.SupplierCode
	record.Address = paramIn.Address
	record.UniformSocialCreditCode = paramIn.UniformSocialCreditCode
	record.Telephone = paramIn.Telephone
	result = util.DB.Debug().Save(&record)
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
	if *paramIn.ChineseName != "" {
		record.ChineseName = paramIn.ChineseName
	}
	if *paramIn.EnglishName != "" {
		record.EnglishName = paramIn.EnglishName
	}
	if *paramIn.SupplierCode != "" {
		record.SupplierCode = paramIn.SupplierCode
	}
	if *paramIn.Address != "" {
		record.Address = paramIn.Address
	}
	if *paramIn.Telephone != "" {
		record.Telephone = paramIn.Telephone
	}
	result := util.DB.Debug().Save(&record)
	if result.Error != nil {
		return serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.Error,
			Message: result.Error.Error(),
		}
	}
	return serializer.ResponseForDetail{
		Data:    nil,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	}
}

func DeleteRelatedParty(id int64) serializer.ResponseForDetail {
	result := util.DB.Debug().Delete(&model.RelatedParty{}, id)
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
