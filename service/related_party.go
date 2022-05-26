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
	ChineseName             string
	EnglishName             string
	SupplierCode            string
	Address                 string
	UniformSocialCreditCode string //统一社会信用代码
	Telephone               string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func RelatedPartyList(paginationRule util.PaginationRule) (*serializer.CommonResponse, error) {
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

func RelatedPartyDetail(id int64) (*serializer.CommonResponse, error) {
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
