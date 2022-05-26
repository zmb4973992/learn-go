package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util"
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
	relatedPartyID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var input service.RelatedParty

	c.ShouldBind(&input)
	input.ID = relatedPartyID
	res := service.UpdateDetailOfRelatedParty(input)
	c.JSON(200, res)
}
