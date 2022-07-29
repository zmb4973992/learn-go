package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/status"
	"net/http"
	"strconv"
)

type ProjectBreakdownController struct {
	baseController
}

func (ProjectBreakdownController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := new(service.ProjectBreakdownService)
	res := s.Get(id)
	c.JSON(http.StatusOK, res)
	return
}
