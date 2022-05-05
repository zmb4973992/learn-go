package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/model"
)

//test
func CreateUser(context *gin.Context) {
	var user model.User
	_ = context.ShouldBind(&user)

}
