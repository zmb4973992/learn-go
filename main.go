package main

import (
	"learn-go/config"
	"learn-go/dao"
	"learn-go/router"
	"learn-go/util/casbin"
	"learn-go/util/logger"
	"learn-go/util/snowflake"
)

func main() {
	//加载配置
	config.Init()
	//加载日志记录器，使用的是zap
	logger.Init()
	//连接数据库
	dao.Init()
	//初始化snowflake，用来生成唯一ID
	snowflake.Init()
	//初始化casbin，用于权限控制
	casbin.Init()
	//开始采用自定义的方式生成引擎
	engine := router.Init()
	err := engine.Run(":" + config.GlobalConfig.HttpPort)
	if err != nil {
		panic(err)
	}
}
