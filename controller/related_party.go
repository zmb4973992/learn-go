package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/dto"
	"learn-go/model"
	"learn-go/serializer"
	"learn-go/service"
	"learn-go/util/status"
	"net/http"
	"strconv"
)

// IRelatedPartyController 使用简单工厂模式,公开接口、公开创建结构体的方法，隐藏结构体
//type IRelatedPartyController interface {
//	Create(c *gin.Context)
//	Get(c *gin.Context)
//	Update(c *gin.Context)
//	Delete(c *gin.Context)
//	List(c *gin.Context)
//}

/* controller层负责接收参数、校验参数
增、改：用model接收    查、删：用id接收     列表：用dto接收，因为有model没有的字段
然后把id或model传给service层进行业务处理
最后拿到service层返回的结果进行展现
*/

// RelatedPartyController 继承baseController，获得相关方法，避免反复重写
type RelatedPartyController struct {
	baseController
}

func NewRelatedPartyController() RelatedPartyController {
	return RelatedPartyController{}
}

func (RelatedPartyController) Get(c *gin.Context) {
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
	c.JSON(http.StatusOK, res)
	return
}

func (RelatedPartyController) Update(c *gin.Context) {
	var paramIn model.RelatedParty
	//先把json参数绑定到model
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
	res := s.Update(&paramIn)
	c.JSON(200, res)
}

func (RelatedPartyController) Create(c *gin.Context) {
	var paramIn model.RelatedParty
	//先把json参数绑定到model
	err := c.ShouldBindJSON(&paramIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidJsonParameters,
			Message: status.GetMessage(status.ErrorInvalidJsonParameters),
		})
		return
	}
	s := service.NewRelatedPartyService()
	res := s.Create(&paramIn)
	c.JSON(http.StatusOK, res)
	return
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

func (RelatedPartyController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ResponseForDetail{
			Data:    nil,
			Code:    status.ErrorInvalidURIParameters,
			Message: status.GetMessage(status.ErrorInvalidURIParameters),
		})
		return
	}
	s := service.NewRelatedPartyService()
	response := s.Delete(id)
	c.JSON(http.StatusOK, response)
}

func (RelatedPartyController) List(c *gin.Context) {
	var relatedPartyListDTO dto.RelatedPartyListDTO
	//这里是bindQuery，只接收query参数
	err := c.ShouldBindQuery(&relatedPartyListDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ResponseForList{
			Data:    nil,
			Paging:  nil,
			Code:    status.ErrorInvalidQueryParameters,
			Message: status.GetMessage(status.ErrorInvalidQueryParameters),
		})
		return
	}
	//生成userService,然后调用它的方法
	s := new(service.RelatedPartyService)
	response := s.List(relatedPartyListDTO)
	c.JSON(http.StatusOK, response)
}
