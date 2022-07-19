package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/serializer"
	"learn-go/util"
	"learn-go/util/status"
	"net/http"
)

func UploadSingle(c *gin.Context) {
	uniqueFilename, err := util.UploadSingleFile(c, "file")
	if err != nil {
		//fmt.Println(err)
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorFailToUploadFiles,
			Message: status.GetMessage(status.ErrorFailToUploadFiles),
		})
		return
	}
	c.JSON(http.StatusOK, serializer.ResponseForDetail{
		Data: gin.H{
			"unique_filename": uniqueFilename,
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	})
	return
}

func UploadMultiple(c *gin.Context) {
	uniqueFilenames, err := util.UploadMultipleFiles(c, "files")
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorFailToUploadFiles,
			Message: status.GetMessage(status.ErrorFailToUploadFiles),
		})
		return
	}
	c.JSON(http.StatusOK, serializer.ResponseForDetail{
		Data: gin.H{
			"unique_filenames": uniqueFilenames,
		},
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	})
	return
}
