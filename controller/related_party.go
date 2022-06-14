package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/dto"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/status"
	"net/http"
	"strconv"
)

// IRelatedPartyController 使用简单工厂模式,公开接口、公开创建结构体的方法，隐藏结构体
type IRelatedPartyController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

//继承baseController，获得相关方法，避免反复重写
type relatedPartyController struct {
	baseController
}

func NewRelatedPartyController() IRelatedPartyController {
	return relatedPartyController{}
}

func (r relatedPartyController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := service.NewRelatedPartyService()
	res := s.Get(id)
	if res == nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorRecordNotFound,
			Message: status.GetMessage(status.ErrorRecordNotFound),
		})
		return
	}
	c.JSON(http.StatusOK, serializer.ResponseForDetail{
		Data:    res,
		Code:    status.Success,
		Message: status.GetMessage(status.Success),
	})
	return
}

func (relatedPartyController) Update(c *gin.Context) {
	var paramIn dto.RelatedPartyDTO
	//先把json参数绑定到dto
	err := c.ShouldBindJSON(&paramIn)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	//把uri上的id参数传递给结构体形式的入参
	paramIn.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := service.NewRelatedPartyService()
	res := s.Update(paramIn)
	c.JSON(200, res)
}

func (relatedPartyController) Create(c *gin.Context) {
	//先声明空的dto，再把context里的数据绑到dto上
	var r dto.RelatedPartyDTO
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	s := service.NewRelatedPartyService()
	res := s.Create(&r)
	c.JSON(http.StatusOK, res)
	return
}

func newd() {

}

//多文件上传
//form, err2 := c.MultipartForm()
//if err2 != nil {
//	fmt.Println(err2)
//	return
//}
//files := form.File["files"]
//for _, file := range files {
//	id := uuid.New().String()
//	file.Filename = id + file.Filename
//	err = c.SaveUploadedFile(file, util.UploadConfig.FullPath+file.Filename)
//	if err != nil {
//		return
//	}
//}
//if err != nil {
//	c.JSON(http.StatusOK, serializer.ResponseForDetail{
//		Data:    nil,
//		Code:    status.Error,
//		Message: status.GetMessage(status.Error),
//	})
//	return
//}

func (relatedPartyController) Delete(c *gin.Context) {
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

func (relatedPartyController) List(c *gin.Context) {
	var s service.RelatedPartyService
	c.ShouldBind(&s)
	//var response serializer.ResponseForList
	response := service.GetRelatedPartyList(s)
	c.JSON(http.StatusOK, response)
}
