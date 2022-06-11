package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"learn-go/config"
	"strings"
)

// UploadSingleFile 上传单个文件专用，经过uuid加持后返回唯一文件名和错误信息。
// 第二个入参为前端的关键词名称。
func UploadSingleFile(c *gin.Context, key string) (uniqueFileName *string, err error) {
	file, err := c.FormFile(key)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	file.Filename = id + "--" + file.Filename
	err = c.SaveUploadedFile(file, config.GlobalConfig.FullPath+file.Filename)
	if err != nil {
		return nil, err
	}
	return &file.Filename, nil
}

// UploadMultipleFiles 上传多个文件专用，经过uuid加持后返回唯一文件名和错误信息，文件名之间用竖线 | 分隔。
// 第二个入参为前端的关键词名称。
func UploadMultipleFiles(c *gin.Context, key string) (uniqueFileNames *string, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	files := form.File[key]
	var fileNames []string
	for _, file := range files {
		id := uuid.New().String()
		file.Filename = id + "--" + file.Filename
		err = c.SaveUploadedFile(file, config.GlobalConfig.FullPath+file.Filename)
		if err != nil {
			return nil, err
		}
		fileNames = append(fileNames, file.Filename)
	}
	res := strings.Join(fileNames, "|")
	return &res, nil
}
