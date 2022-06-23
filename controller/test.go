package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/util"
	"net/http"
)

func Upload(c *gin.Context) {
	res, _ := util.UploadSingleFile(c, "file")
	c.JSON(http.StatusOK, gin.H{
		"name": res,
	})

}

func UploadMutiple(c *gin.Context) {
	res, _ := util.UploadMultipleFiles(c, "files")
	c.JSON(http.StatusOK, gin.H{
		"name": res,
	})
}
