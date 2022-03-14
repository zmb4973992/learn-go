package model

type RelatedParty struct {
	ID           uint
	ChineseName  string
	EnglishName  string
	Relationship string
	SupplierCode string
	//Address                 string
	//UniformSocialCreditCode string //统一社会信用代码
	//Telephone               string
	//CreatedAt               time.Time
	//UpdatedAt               time.Time
	//Project []Project `gorm:"foreignkey:RelatedPartyID;constraint:OnUpdate:CASCADE;"`
}

//func (RelatedParty) TableName() string {
//	return "related_party"
//}
