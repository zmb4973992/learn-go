package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/dto"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/status"
	"net/http"
	"strconv"
)

type DepartmentController struct {
	baseController
}

func NewDepartmentController() DepartmentController {
	return DepartmentController{}
}

func (DepartmentController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := service.NewDepartmentService()
	res := s.Get(id)
	c.JSON(http.StatusOK, res)
	return
}

func (DepartmentController) Create(c *gin.Context) {
	var param dto.DepartmentCreateAndUpdateDTO
	//先把json参数绑定到model
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	s := service.NewDepartmentService()
	res := s.Create(&param)
	c.JSON(http.StatusOK, res)
	return
}

func (DepartmentController) Update(c *gin.Context) {
	var param dto.DepartmentCreateAndUpdateDTO
	//先把json参数绑定到model
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	//把uri上的id参数传递给结构体形式的入参
	param.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := service.NewDepartmentService()
	res := s.Update(&param)
	c.JSON(200, res)
}

func (DepartmentController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := service.NewDepartmentService()
	response := s.Delete(id)
	c.JSON(http.StatusOK, response)
}
