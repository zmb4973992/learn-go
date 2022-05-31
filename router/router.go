package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
)

// InitRouter 初始化路由器,最终返回*gin.Engine类型，给main调用
func InitRouter() *gin.Engine {
	//使用gin框架，生成默认的空引擎
	engine := gin.Default()
	engine.MaxMultipartMemory = 512 << 20 //512MB
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
		user := api.Group("/user")
		{
			user.GET("/list", controller.GetUserList)  //获取用户列表
			user.GET("/:id", controller.GetUser)       //获取用户详情
			user.PUT("/:id", controller.UpdateUser)    //修改用户
			user.POST("/", controller.CreateUser)      //添加用户
			user.DELETE("/:id", controller.DeleteUser) //删除用户
		}
		relatedParty := api.Group("/related_party")
		{
			relatedParty.GET("/list", controller.GetRelatedPartyList)  //获取列表
			relatedParty.GET("/:id", controller.GetRelatedParty)       //获取详情
			relatedParty.PUT("/:id", controller.UpdateRelatedParty)    //修改详情
			relatedParty.POST("", controller.CreateRelatedParty)       //添加相关方详情
			relatedParty.DELETE("/:id", controller.DeleteRelatedParty) //删除详情
		}

	}

	//引擎处理完成后，返回
	return engine
}
