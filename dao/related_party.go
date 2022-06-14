package dao

import (
	"learn-go/dto"
	"learn-go/model"
)

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

func (RelatedPartyDAO) Update(id int, paramIn *model.RelatedParty) error {
	err := DB.Debug().Where("id = ?", id).Updates(&paramIn).Error
	return err
}
