package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/model"
	"learn-go/util"
)

func Test(c *gin.Context) {
	var x model.Test
	c.ShouldBind(&x)

	util.DB.Create(&x)
}
