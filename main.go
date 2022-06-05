package main

import (
	"learn-go/router"
	"learn-go/util"
	"learn-go/util/logger"
	"learn-go/util/snowflake"
)

func main() {
	//加载配置
	util.LoadConfig()
	//加载日志记录器，使用的是zap
	logger.InitLogger()
	//连接数据库
	util.ConnectDB()
	//初始化snowflake，用来生成唯一ID
	err := snowflake.InitSnowFlake()
	if err != nil {
		panic("生成snowflake实例失败，请重试")
	}

	//开始采用自定义的方式生成引擎
	engine := router.InitRouter()
	err = engine.Run(":" + util.GeneralConfig.HttpPort)
	if err != nil {
		panic(err)
	}
}
