package dto

import (
	"time"
)

type relatedPartyDTO struct {
	ID                      int
	ChineseName             *string    `json:"chinese_name"`
	EnglishName             *string    `json:"english_name" `
	SupplierCode            *string    `json:"supplier_code" `
	Address                 *string    `json:"address" `
	UniformSocialCreditCode *string    `json:"uniform_social_credit_code" ` //统一社会信用代码
	Telephone               *string    `json:"telephone" `
	File                    *string    `json:"-"`
	Files                   *string    `json:"-"`
	CreatedAt               *time.Time `json:"created_at"`
	UpdatedAt               *time.Time `json:"updated_at"`
	Paging                  PagingDTO  `json:"-"`
}

func (r relatedPartyDTO) Get(int) {
	//TODO implement me
	panic("implement me")
}

func (r relatedPartyDTO) GetList() {
	//TODO implement me
	panic("implement me")
}

func (r relatedPartyDTO) Update() {
	//TODO implement me
	panic("implement me")
}

func (r relatedPartyDTO) Create() {
	//TODO implement me
	panic("implement me")
}

func (r relatedPartyDTO) Delete() {
	//TODO implement me
	panic("implement me")
}
