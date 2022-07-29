package dto

type ProjectBreakdownGetDTO struct {
	Name      *string  `json:"name"`       //名称
	ProjectID *int     `json:"project_id"` //所属项目id
	Level     *int     `json:"level"`      //层级
	Weight    *float64 `json:"weight"`     //权重
	Superior  any      `json:"superior"`   //上级信息
}

type ProjectBreakdownCreateAndUpdateDTO struct {
	ID         int      `json:"id"`
	Name       *string  `json:"name" binding:"required"`        //拆解项名称
	ProjectID  *int     `json:"project_id" binding:"required"`  //所属项目id
	Level      *int     `json:"level" binding:"required"`       //层级
	Weight     *float64 `json:"weight" binding:"required"`      //权重
	SuperiorID *int     `json:"superior_id" binding:"required"` //上级拆解项ID
}
