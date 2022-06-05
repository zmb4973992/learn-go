package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
	"learn-go/middleware"
	"learn-go/util/logger"
)

// InitRouter 初始化路由器,最终返回*gin.Engine类型，给main调用
func InitRouter() *gin.Engine {
	//使用gin框架，生成默认的空引擎
	engine := gin.New()
	engine.Use(logger.ZapLogger(logger.ZapStandardLogger), gin.Recovery())
	engine.POST("/test", controller.Test)
	engine.Use(middleware.Cors())
	engine.POST("/login", controller.Login)         //用户登录
	engine.POST("/api/user", controller.CreateUser) //添加用户
	//依次加载所有的路由组
	api := engine.Group("/api")
	{
		//api下都需要登录后操作
		api.Use(middleware.JWT())
		{
			user := api.Group("/user")
			{
				user.GET("/list", controller.GetUserList)  //获取用户列表
				user.GET("/:id", controller.GetUser)       //获取用户详情
				user.PUT("/:id", controller.UpdateUser)    //修改用户
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
	}

	//引擎处理完成后，返回
	return engine
}