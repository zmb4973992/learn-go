package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/model"
)

func CreateUser(context *gin.Context) {
	var user model.User
	_ = context.ShouldBind(&user)
	fmt.Println(user)
	fmt.Println(user.Username)
}
