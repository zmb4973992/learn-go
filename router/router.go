package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
	"learn-go/middleware"
	"learn-go/util/logger"
)

// Init 初始化路由器,最终返回*gin.Engine类型，给main调用
func Init() *gin.Engine {
	//使用gin框架，生成默认的空引擎
	engine := gin.New()
	//使用中间件
	engine.Use(logger.ZapLogger(), gin.Recovery())
	engine.Use(middleware.Cors())
	engine.POST("/api/login", controller.Login) //用户登录

	//创建所有的控制器
	userController := controller.NewUserController()
	relatedPartyController := controller.NewRelatedPartyController()

	engine.POST("/api/user", userController.Create) //添加用户
	//依次加载所有的路由组，以下都需要经过jwt验证
	api := engine.Group("/api").Use(middleware.JWT())
	{
		api.GET("/user/:id", userController.Get)       //获取用户详情
		api.PUT("/user/:id", userController.Update)    //修改用户
		api.DELETE("/user/:id", userController.Delete) //删除用户
		api.GET("/user/list", userController.List)     //获取用户列表

		api.GET("/related_party/list", relatedPartyController.GetRelatedPartyList) //获取列表
		api.GET("/related_party/:id", relatedPartyController.GetRelatedParty)      //获取详情
		api.PUT("/related_party/:id", relatedPartyController.UpdateRelatedParty)   //修改详情
		api.POST("/related_party", relatedPartyController.CreateRelatedParty)      //添加相关方详情
		api.DELETE("/:id", relatedPartyController.DeleteRelatedParty)              //删除详情
	}

	//引擎处理完成后，返回
	return engine
}
