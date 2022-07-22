package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/model"
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
	var paramIn model.Department
	//先把json参数绑定到model
	err := c.ShouldBindJSON(&paramIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	s := service.NewDepartmentService()
	res := s.Create(&paramIn)
	c.JSON(http.StatusOK, res)
	return
}
