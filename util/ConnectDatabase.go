package util

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"learn-go/model"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func test() {

}

//待重构测试~
func ConnectDatabase() {
	//通过gorm连接sqlserver数据库
	DB, err = gorm.Open(sqlserver.Open(MyConfig.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//使用gorm标准格式，创建连接池
	sqlDB, err := DB.DB()
	// Set Max Idle Connections 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// Set Max Open Connections 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// Set Connection Max Lifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = DB.AutoMigrate(
		&model.RelatedParty{},
		&model.Project{},
		&model.User{},
	)
	if err != nil {
		fmt.Println("数据表迁移失败，请检查")
	}
}
