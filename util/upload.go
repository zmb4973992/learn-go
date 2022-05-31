package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

// UploadSingleFile 上传单个文件专用，经过uuid加持后返回唯一文件名和错误信息。
// 第二个入参为前端的关键词名称。
func UploadSingleFile(c *gin.Context, key string) (uniqueFileName *string) {
	file, err := c.FormFile(key)
	if err != nil {
		return nil
	}
	if file.Size > MyUploadConfig.MaxSize {
		return nil
	}
	id := uuid.New().String()
	file.Filename = id + "--" + file.Filename
	err = c.SaveUploadedFile(file, MyUploadConfig.FullPath+file.Filename)
	if err != nil {
		log.Print(err)
		return nil
	}
	return &file.Filename
}

// UploadMultipleFiles 上传多个文件专用，经过uuid加持后返回唯一文件名和错误信息，文件名之间用竖线 | 分隔。
// 第二个入参为前端的关键词名称。
func UploadMultipleFiles(c *gin.Context, key string) (uniqueFileNames *string) {
	form, _ := c.MultipartForm()
	files := form.File[key]
	var fileNames []string
	for _, file := range files {
		log.Println(file.Filename)
		fileNames = append(fileNames, file.Filename)
		log.Println(fileNames)
	}
	return nil
}
