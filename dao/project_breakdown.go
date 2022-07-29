package dao

import (
	"learn-go/dto"
	"learn-go/model"
	"strconv"
)

type ProjectBreakdownDAO struct{}

func (ProjectBreakdownDAO) Get(projectBreakdownID int) *dto.ProjectBreakdownGetDTO {
	//默认嵌套递归次数上限为4次，太多了降低效率，而且没必要
	return getProjectBreakdownWithRecursionLimit(projectBreakdownID, 4, 0)
}

//由于get方法有递归调用，所以需要在这里多加2个参数进行限制。标准的get方法调用这个内部函数，达到封装的效果
func getProjectBreakdownWithRecursionLimit(projectBreakdownID int, recursionTimesLimit int, recursionTimes int) *dto.ProjectBreakdownGetDTO {
	var projectBreakdownGetDTO = dto.ProjectBreakdownGetDTO{}
	//把基础的拆解信息查出来
	var projectBreakdown model.ProjectBreakdown
	err := model.DB.Where("id = ?", projectBreakdownID).First(&projectBreakdown).Error
	if err != nil {
		return nil
	}
	//把所有查出的结果赋值给输出变量
	projectBreakdownGetDTO.Name = projectBreakdown.Name
	projectBreakdownGetDTO.ProjectID = projectBreakdown.ProjectID
	projectBreakdownGetDTO.Level = projectBreakdown.Level
	projectBreakdownGetDTO.Weight = projectBreakdown.Weight

	//递归查询上级信息
	if projectBreakdown.SuperiorID != nil {
		recursionTimes += 1
		if recursionTimes <= recursionTimesLimit {
			projectBreakdownGetDTO.Superior = getProjectBreakdownWithRecursionLimit(*projectBreakdown.SuperiorID, recursionTimesLimit, recursionTimes)
		} else {
			projectBreakdownGetDTO.Superior = "递归深度超过" + strconv.Itoa(recursionTimesLimit) + "次，可能存在循环递归，请检查数据是否正确"
		}
	}
	return &projectBreakdownGetDTO
}

// Create 这里是只负责新增，不写任何业务逻辑。只要收到参数就创建数据库记录，然后返回错误
func (ProjectBreakdownDAO) Create(param *model.ProjectBreakdown) error {
	err := model.DB.Create(param).Error
	return err
}
