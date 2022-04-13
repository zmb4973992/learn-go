package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
)

func LoadUserRouter(engine *gin.Engine) {
	UserGroup := engine.Group("/user")
	{
		UserGroup.POST("/1", controller.CreateUser)
	}

}
