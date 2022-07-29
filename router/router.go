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
	userController := new(controller.UserController)
	relatedPartyController := new(controller.RelatedPartyController)
	noRouteController := new(controller.NoRouteController)
	departmentController := new(controller.DepartmentController)
	projectBreakdownController := new(controller.ProjectBreakdownController)

	engine.POST("/api/user", userController.Create)            //添加用户
	engine.POST("/upload_single", controller.UploadSingle)     //测试上传单个
	engine.POST("/upload_multiple", controller.UploadMultiple) //测试上传多个
	//依次加载所有的路由组，以下都需要登录验证(jwt验证)
	api := engine.Group("/api").Use(middleware.NeedLogin())
	{
		api.GET("/user/:id", middleware.NeedAuth(), userController.Get) //获取用户详情
		api.PUT("/user/:id", userController.Update)                     //修改用户
		api.DELETE("/user/:id", userController.Delete)                  //删除用户
		api.GET("/user/list", userController.List)                      //获取用户列表

		api.GET("/related_party/list", relatedPartyController.List)  //获取相关方列表
		api.GET("/related_party/:id", relatedPartyController.Get)    //获取相关方详情
		api.PUT("/related_party/:id", relatedPartyController.Update) //修改相关方
		api.POST("/related_party", relatedPartyController.Create)    //新增相关方
		api.DELETE("/:id", relatedPartyController.Delete)            //删除相关方

		api.GET("/department/:id", departmentController.Get)       //获取部门详情
		api.POST("/department", departmentController.Create)       //新增部门
		api.PUT("/department/:id", departmentController.Update)    //修改部门
		api.DELETE("/department/:id", departmentController.Delete) //删除部门

		api.GET("/project_breakdown/:id", projectBreakdownController.Get) //获取项目拆解详情
	}

	engine.NoRoute(noRouteController.NoRoute)

	//引擎处理完成后，返回
	return engine
}
