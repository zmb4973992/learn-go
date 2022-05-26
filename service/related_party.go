package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/code"
	"time"
)

type RelatedParty struct {
	ID                      int64
	ChineseName             *string   `form:"chinese_name"`
	EnglishName             *string   `form:"english_name"`
	SupplierCode            *string   `form:"supplier_code"`
	Address                 string    `form:"address"`
	UniformSocialCreditCode string    `form:"uniform_social_credit_code"` //统一社会信用代码
	Telephone               string    `form:"telephone"`
	CreatedAt               time.Time `form:"created_at"`
	UpdatedAt               time.Time `form:"updated_at"`
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
			Code:    code.Error_Record_Not_Found,
			Message: code.GetErrorMessage(code.Error_Record_Not_Found),
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
			Code:    code.Error_Record_Not_Found,
			Message: code.GetErrorMessage(code.Error_Record_Not_Found),
		}
	}
	record.ChineseName = paramIn.ChineseName
	record.EnglishName = paramIn.EnglishName
	result = util.DB.Debug().Save(&record)
	if result.Error != nil {
		fmt.Println(result.Error)
		return serializer.CommonResponse{}
	}
	return serializer.CommonResponse{
		Data:    nil,
		Code:    code.Success,
		Message: code.GetErrorMessage(code.Success),
	}

}
