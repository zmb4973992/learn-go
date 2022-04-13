package router

import (
	"github.com/gin-gonic/gin"
)

//初始化路由器,最终返回*gin.Engine类型，给main调用

func InitRouter() *gin.Engine {
	//使用gin框架，生成默认的空引擎
	engine := gin.Default()
	//挂载自定义中间件
	//engine.Use(middleware.MyMiddlewareTest())
	//依次加载所有的路由组
	LoadRelatedPartyRouter(engine)
	LoadProjectRouter(engine)
	LoadUserRouter(engine)
	//引擎处理完成后，返回
	return engine
}
