package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util"
	"net/http"
	"strconv"
)

func RelatedPartyDetail(c *gin.Context) {
	relatedPartyID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res, _ := service.RelatedPartyDetail(relatedPartyID)
	c.JSON(http.StatusOK, res)
}

func RelatedPartyList(c *gin.Context) {
	var paginationRule util.PaginationRule
	c.ShouldBind(&paginationRule)
	var response *serializer.CommonResponse
	response, _ = service.RelatedPartyList(paginationRule)
	c.JSON(http.StatusOK, response)
}
