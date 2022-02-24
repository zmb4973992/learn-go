package model

type User struct {
	ID        uint
	Name      string
	Telephone string
	Mycard    []CreditCard `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreditCard struct {
	Number int
	UserID uint //gorm自动将该字段定义为外键
}

func (CreditCard) TableName() string {
	return "mycard"
}
