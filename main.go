package main

import (
	"learn-go/router"
	"learn-go/util"
)

//test
func main() {
	//加载配置
	util.LoadConfig()
	//连接数据库
	util.ConnectDatabase()
	//开始采用自定义的方式生成引擎
	engine := router.InitRouter()
	err := engine.Run(":" + util.MyConfig.HttpPort)
	if err != nil {
		panic("端口错误或未知错误...")
	}
}
