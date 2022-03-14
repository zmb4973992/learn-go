package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
)

func LoadProjectRouter(engine *gin.Engine) {
	ProjectGroup := engine.Group("/project")
	{
		ProjectGroup.GET("/1", controller.Handler2)
	}

}
