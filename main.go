package main

import (
	"learn-go/router"
	"learn-go/util"
	"learn-go/util/logging"
	"log"
)

//test
func main() {
	log.Print("111")

	//加载配置
	util.LoadConfig()
	//加载日志记录器，使用的是zap
	logging.InitLogger()
	//连接数据库
	util.ConnectDB()

	//开始采用自定义的方式生成引擎
	engine := router.InitRouter()
	err := engine.Run(":" + util.MyConfig.HttpPort)
	if err != nil {
		panic(err)
	}

}
