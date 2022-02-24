package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
)

func LoadGroup1Router(engine *gin.Engine) {
	Group1 := engine.Group("/group1")
	{
		Group1.GET("/1", controller.Handler1)
	}

}
