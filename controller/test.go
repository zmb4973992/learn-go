package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/dao"
	"learn-go/model"
)

func Test(c *gin.Context) {
	var x model.Test
	c.ShouldBind(&x)

	dao.DB.Create(&x)
}
