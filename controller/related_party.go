package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util"
	"learn-go/util/status"
	"net/http"
	"strconv"
	"strings"
)

func GetRelatedPartyList(c *gin.Context) {
	var paginationRule util.PagingRule
	c.ShouldBind(&paginationRule)
	var response serializer.ResponseForList
	response = service.GetRelatedPartyList(paginationRule)
	c.JSON(http.StatusOK, response)
}

func GetRelatedParty(c *gin.Context) {
	relatedPartyID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res, _ := service.GetDetailOfRelatedParty(relatedPartyID)
	c.JSON(http.StatusOK, res)

}

func UpdateRelatedParty(c *gin.Context) {
	var paramIn service.RelatedPartyService
	err := c.ShouldBind(&paramIn)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.Error,
			Message: status.GetMessage(status.Error),
		})
		return
	}
	paramIn.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64) //把uri上的id参数传递给结构体形式的入参
	res := service.UpdateDetailOfRelatedParty(paramIn)
	c.JSON(200, res)
}

func CreateRelatedParty(c *gin.Context) {
	var paramIn service.RelatedPartyService
	err := c.ShouldBind(&paramIn)
	//单文件上传
	file1, err1 := c.FormFile("file1")
	b := strings.Join(file1.Filename, "|")
	if err1 != nil {
	} else {
		id := uuid.New().String()
		file1.Filename = id + file1.Filename
		err = c.SaveUploadedFile(file1, util.MyUploadConfig.FullPath+file1.Filename)
		if err != nil {
			return
		}
	}

	//多文件上传
	form, err2 := c.MultipartForm()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	files := form.File["files"]
	for _, file := range files {
		id := uuid.New().String()
		file.Filename = id + file.Filename
		err = c.SaveUploadedFile(file, util.MyUploadConfig.FullPath+file.Filename)
		if err != nil {
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.Error,
			Message: status.GetMessage(status.Error),
		})
		return
	}
	res := service.CreateRelatedParty(paramIn)
	c.JSON(http.StatusOK, res)
}

func DeleteRelatedParty(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	res := service.DeleteRelatedParty(id)
	c.JSON(http.StatusOK, res)
}
