package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/dto"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/status"
	"net/http"
	"strconv"
)

type IUserController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type userController struct {
	baseController
}

func NewUserController() IUserController {
	return userController{}
}

func (userController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := service.NewUserService()
	res := s.Get(id)
	c.JSON(http.StatusOK, res)
	return
}

func (userController) Create(c *gin.Context) {
	//先声明空的dto，再把context里的数据绑到dto上
	var u dto.UserDTO
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	s := service.NewUserService()
	res := s.Create(&u)
	c.JSON(http.StatusOK, res)
	return
}

// Update controller的功能：解析uri参数、json参数，拦截非法参数，然后传给service层处理
func (userController) Update(c *gin.Context) {
	//这里只更新传过来的参数，所以采用map形式
	paramIn := make(map[string]any)
	_ = c.ShouldBindJSON(&paramIn)
	//把uri上的id参数传递给结构体形式的入参
	id, err := strconv.Atoi(c.Param("id"))
	//如果解析失败，例如URI的参数不是数字
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := new(service.UserService)
	//参数解析完毕，交给service层处理
	res := s.Update(id, paramIn)
	c.JSON(200, res)
}

func (userController) Delete(c *gin.Context) {
	//把uri上的id参数传递给结构体形式的入参
	id, err := strconv.Atoi(c.Param("id"))
	//如果解析失败，例如URI的参数不是数字
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := new(service.UserService)
	response := s.Delete(id)
	c.JSON(200, response)
}

func (userController) List(c *gin.Context) {
	type a struct {
		Fields    []string `json:"fields"`
		Ascending bool     `json:"ascending"`
	}
	var b a
	c.ShouldBindJSON(&b)
	fmt.Println(b.Fields)
	//这里只处理传过来的参数，所以采用map形式,打包传给service层进行处理
	paramIn := make(map[string]any)
	paramIn["fields"] = []string{}
	err := c.ShouldBindJSON(&paramIn)

	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	s := new(service.UserService)
	response := s.List(paramIn)
	c.JSON(http.StatusOK, response)
}
