package dao

import (
	"learn-go/dto"
	"learn-go/model"
)

/*
dao层的基本原则：
入参为id或model，用于对数据库进行增删改查；
出参为err或dto，用于反馈结果或给其他层使用
*/

func NewRelatedDAO() RelatedPartyDAO {
	return RelatedPartyDAO{}
}

// RelatedPartyDAO dao层的结构体没有数据，只是操作数据库进行增删改查，不写业务逻辑
type RelatedPartyDAO struct{}

func (RelatedPartyDAO) Get(id int) *dto.RelatedPartyDTO {
	//之所以用dto不用model，是因为model为数据库原表，数据可能包含敏感字段、或未加工，不适合直接传递
	//传递的功能基本都交给dto
	var r dto.RelatedPartyDTO
	err := DB.Model(&model.RelatedParty{}).Debug().Where("id = ?", id).First(&r).Error
	if err != nil {
		return nil
	}
	return &r
}

// Create 这里是只负责新增，不写任何业务逻辑。
// 创建数据库记录，返回错误
func (RelatedPartyDAO) Create(paramIn *model.RelatedParty) error {
	err := DB.Debug().Create(&paramIn).Error
	return err
}

func (RelatedPartyDAO) Update(paramIn *model.RelatedParty) error {
	err = DB.Debug().Where("id = ?", paramIn.ID).Updates(paramIn).Error
	if *paramIn.ChineseName == "<empty>" {

	}
	return err
}

func (RelatedPartyDAO) Delete(id int) error {
	err := DB.Debug().Delete(&model.RelatedParty{}, id).Error

	return err
}
