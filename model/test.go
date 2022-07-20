package model

type Test struct {
	ID   int64
	Name string
	//外键
	ProjectID int
}

// TableName 将表名改为project
func (Test) TableName() string {
	return "test"
}
