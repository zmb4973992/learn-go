package service

import (
	"errors"
	"gorm.io/gorm"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/code"
	"time"
)

type RelatedParty struct {
	ID                      int64
	ChineseName             *string    `form:"chinese_name" binding:"required"`
	EnglishName             *string    `form:"english_name" binding:"required"`
	SupplierCode            *string    `form:"supplier_code" binding:"required"`
	Address                 *string    `form:"address" binding:"required"`
	UniformSocialCreditCode *string    `form:"uniform_social_credit_code" binding:"required"` //统一社会信用代码
	Telephone               *string    `form:"telephone" binding:"required"`
	CreatedAt               *time.Time `form:"created_at"`
	UpdatedAt               *time.Time `form:"updated_at"`
}

func GetListOfRelatedParty(paginationRule util.PaginationRule) (*serializer.CommonResponse, error) {
	var list []*model.RelatedParty
	if paginationRule.Page <= 0 {
		paginationRule.Page = 1
	}
	if paginationRule.PageSize <= 0 || paginationRule.PageSize > 100 {
		paginationRule.PageSize = 20
	}
	util.DB.Debug().Scopes(util.Paginate(&paginationRule)).Find(&list)
	return &serializer.CommonResponse{
		Data: list,
	}, nil
}

func GetDetailOfRelatedParty(id int64) (*serializer.CommonResponse, error) {
	var record *model.RelatedParty
	result := util.DB.Debug().Where("id=?", id).First(&record)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &serializer.CommonResponse{
			Data:    nil,
			Code:    code.ErrorRecordNotFound,
			Message: code.GetErrorMessage(code.ErrorRecordNotFound),
		}, nil
	}
	return &serializer.CommonResponse{
		Data:    record,
		Code:    code.Success,
		Message: code.GetErrorMessage(code.Success),
	}, nil
}

func UpdateDetailOfRelatedParty(paramIn RelatedParty) serializer.CommonResponse {
	var record model.RelatedParty
	result := util.DB.Debug().First(&record, paramIn.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return serializer.CommonResponse{
			Data:    nil,
			Code:    code.ErrorRecordNotFound,
			Message: code.GetErrorMessage(code.ErrorRecordNotFound),
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
		return serializer.CommonResponse{
			Data:    nil,
			Code:    code.Error,
			Message: code.GetErrorMessage(code.Error),
		}
	}
	return serializer.CommonResponse{
		Data:    nil,
		Code:    code.Success,
		Message: code.GetErrorMessage(code.Success),
	}
}

func CreateRelatedParty(paramIn RelatedParty) serializer.CommonResponse {
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
		return serializer.CommonResponse{
			Data:    nil,
			Code:    code.Error,
			Message: result.Error.Error(),
		}
	}
	return serializer.CommonResponse{
		Data:    nil,
		Code:    code.Success,
		Message: code.GetErrorMessage(code.Success),
	}
}

func DeleteRelatedParty(id int64) serializer.CommonResponse {
	result := util.DB.Debug().Delete(&model.RelatedParty{}, id)
	if result.Error != nil {
		return serializer.CommonResponse{
			Data:    nil,
			Code:    code.Error,
			Message: result.Error.Error(),
		}
	}
	return serializer.CommonResponse{
		Data:    nil,
		Code:    code.Success,
		Message: code.GetErrorMessage(code.Success),
	}
}
