package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"learn-go/model"
	"learn-go/router"
	"learn-go/util"
	"time"
)

func main() {

	//加载配置
	util.LoadConfig()
	//通过gorm连接sqlserver数据库
	db, _ := gorm.Open(sqlserver.Open(util.MyConfig.Dsn), &gorm.Config{})
	dbPool, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbPool.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	dbPool.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	dbPool.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{}, &model.CreditCard{})
	//开始采用自定义的方式生成引擎
	engine := router.InitRouter()
	engine.Run(":" + util.MyConfig.HttpPort)

}
