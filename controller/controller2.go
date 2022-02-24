package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//所有controller都放在这里，基本上是一个routerGroup对应一个controller文件

func Handler2(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "This is controller 2",
	})
}
