package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
)

// InitRouter 初始化路由器,最终返回*gin.Engine类型，给main调用
func InitRouter() *gin.Engine {
	//使用gin框架，生成默认的空引擎
	engine := gin.Default()
	//挂载自定义中间件
	//engine.Use(middleware.MyMiddlewareTest())
	//依次加载所有的路由组
	api := engine.Group("/api/v1")
	{
		api.POST("/login", controller.Login)

		//api.Use(middleware.JWT())
		{
			api.POST("/test", controller.ParseTokenTest)

		}
		api.GET("/user/:id")
		relatedParty := api.Group("/related_party")
		{
			relatedParty.GET("/list", controller.RelatedPartyList)
			relatedParty.GET("/:id", controller.RelatedPartyDetail)
		}

	}

	//引擎处理完成后，返回
	return engine
}
