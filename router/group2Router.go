package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
)

func LoadGroup2Router(engine *gin.Engine) {
	Group2 := engine.Group("/group2")
	{
		Group2.GET("/1", controller.Handler2)
	}

}
