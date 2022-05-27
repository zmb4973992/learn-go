package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/code"
	"net/http"
	"strconv"
)

// CreateUser 测试
func CreateUser(context *gin.Context) {
	var user model.User
	_ = context.ShouldBind(&user)
}

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, serializer.CommonResponse{
			Data:    nil,
			Code:    code.ErrorInvalidParameters,
			Message: code.GetErrorMessage(code.ErrorInvalidParameters),
		})
		return
	}
	res := service.GetUser(id)
	c.JSON(http.StatusOK, res)
}
