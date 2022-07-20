package model

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"learn-go/config"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	//通过gorm连接sqlserver数据库
	DB, err = gorm.Open(sqlserver.Open(config.GlobalConfig.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//使用gorm标准格式，创建连接池
	sqlDB, _ := DB.DB()
	// Set Max Idle Connections 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// Set Max Open Connections 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// Set Connection Max Lifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = DB.AutoMigrate(
		&RelatedParty{},
		&Project{},
		&Department{},
		&User{},
		&DepartmentAndUser{},
		&Role{},
		&RoleAndUser{},
	)

	if err != nil {
		panic(err)
	}
}

func GetTotalPages(totalRecords int, pageSize int) (totalPages int) {
	if totalRecords <= 0 || pageSize <= 0 {
		return 0
	}
	totalPages = totalRecords / pageSize
	if totalRecords%pageSize != 0 {
		totalPages++
	}
	return totalPages
}
