package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/util"
	"net/http"
)

func UploadSingle(c *gin.Context) {
	uniqueFilename, err := util.UploadSingleFile(c, "file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "上传失败，请重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":          "ok",
		"unique_filename": uniqueFilename,
	})
	return
}

func UploadMultiple(c *gin.Context) {
	uniqueFilenames, err := util.UploadMultipleFiles(c, "files")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "上传失败，请重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":           "ok",
		"unique_filenames": uniqueFilenames,
	})
	return
}
