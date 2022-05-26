package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util"
	"learn-go/util/code"
	"net/http"
	"strconv"
)

func GetListOfRelatedParty(c *gin.Context) {
	var paginationRule util.PaginationRule
	c.ShouldBind(&paginationRule)
	var response *serializer.CommonResponse
	response, _ = service.GetListOfRelatedParty(paginationRule)
	c.JSON(http.StatusOK, response)
}

func GetDetailOfRelatedParty(c *gin.Context) {
	relatedPartyID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res, _ := service.GetDetailOfRelatedParty(relatedPartyID)
	c.JSON(http.StatusOK, res)
}

func UpdateDetailOfRelatedParty(c *gin.Context) {
	var paramIn service.RelatedParty
	err := c.ShouldBind(&paramIn)
	if err != nil {
		c.JSON(http.StatusOK, serializer.CommonResponse{
			Data:    nil,
			Code:    code.ErrorNotEnoughParameters,
			Message: code.GetErrorMessage(code.ErrorNotEnoughParameters),
		})
		return
	}
	paramIn.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64) //把uri上的id参数传递给结构体形式的入参
	res := service.UpdateDetailOfRelatedParty(paramIn)
	c.JSON(200, res)
}
