package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/status"
	"net/http"
	"strconv"
)

type IUserController interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserList(c *gin.Context)
}

type userController struct {
	baseController
}

func NewUserController() IUserController {
	return userController{}
}

func (userController) CreateUser(c *gin.Context) {
	var record service.UserService
	err := c.ShouldBind(&record)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidFormDataParameters,
			Message: status.GetMessage(status.ErrorInvalidFormDataParameters),
		})
		return
	}
	res := record.Create(record)
	c.JSON(http.StatusOK, res)
	return
}

func (userController) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	res := service.GetUser(id)
	c.JSON(http.StatusOK, res)
	return
}

func (userController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //把uri上的id参数传递给结构体形式的入参
	//如果URI的参数不是数字
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	var paramIn service.UserService
	paramIn.ID = id
	err = c.ShouldBind(&paramIn)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidFormDataParameters,
			Message: status.GetMessage(status.ErrorInvalidFormDataParameters),
		})
		return
	}
	res := service.UpdateUser(paramIn)
	c.JSON(200, res)
}

func (userController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64) //把uri上的id参数传递给结构体形式的入参
	//如果URI的参数不是数字
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	res := service.DeleteUser(id)
	c.JSON(200, res)
}

func (userController) GetUserList(c *gin.Context) {
	var paginationRule service.Paging
	err := c.ShouldBind(&paginationRule) //不需要处理错误，如果绑定不上，下面的方法会自动使用默认值
	if err != nil {
	}
	var response serializer.ResponseForDetail
	response = service.GetUserList(paginationRule)
	c.JSON(http.StatusOK, response)
}
