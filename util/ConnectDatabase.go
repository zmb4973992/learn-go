package util

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"learn-go/model"
	"time"
)

func ConnectDatabase() {
	//通过gorm连接sqlserver数据库
	db, _ := gorm.Open(sqlserver.Open(MyConfig.Dsn), &gorm.Config{})

	//使用gorm标准格式，创建连接池
	dbPool, _ := db.DB()
	// Set Max Idle Connections 设置空闲连接池中连接的最大数量
	dbPool.SetMaxIdleConns(10)

	// Set Max Open Connections 设置打开数据库连接的最大数量
	dbPool.SetMaxOpenConns(100)

	// Set Connection Max Lifetime 设置了连接可复用的最大时间
	dbPool.SetConnMaxLifetime(time.Hour)

	err := db.Debug().AutoMigrate(&model.RelatedParty{}, &model.Project{})
	if err != nil {
		fmt.Println("数据表创建失败，请检查")
	}
}
