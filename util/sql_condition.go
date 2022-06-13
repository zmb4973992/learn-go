package util

import (
	"gorm.io/gorm"
)

type SqlCondition struct {
	SelectedColumns []string
	ParamPairs      []ParamPair
	OrderBy         OrderBy
	Paging          Paging
}

type ParamPair struct {
	ParamKey   string //查询参数的名称，如 age>=, name include, id=
	ParamValue any    //查询参数的值
}

type OrderBy struct {
	Column    string //排序字段
	Ascending bool   //是否为升序（从小到大）
}

type Paging struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// NewSqlCondition 生成自定义的查询条件,参数可不填
//必须为指针，因为下面的方法要用到指针进行修改入参
func NewSqlCondition() *SqlCondition {
	return &SqlCondition{}
}

// Where 给SqlCondition自定义where方法，将参数保存到ParameterPair中
func (s *SqlCondition) Where(key string, value any) *SqlCondition {
	s.ParamPairs = append(s.ParamPairs, ParamPair{
		ParamKey:   key,
		ParamValue: value,
	})
	return s
}

func (s *SqlCondition) Equal(parameterKey string, parameterValue any) *SqlCondition {
	s.Where(parameterKey+" = ? ", parameterValue)
	return s
}

func (s *SqlCondition) NotEqual(parameterKey string, parameterValue any) *SqlCondition {
	s.Where(parameterKey+" <> ?", parameterValue)
	return s
}

func (s *SqlCondition) Gt(parameterKey string, parameterValue int) *SqlCondition {
	s.Where(parameterKey+" > ?", parameterValue)
	return s
}

func (s *SqlCondition) Gte(parameterKey string, parameterValue int) *SqlCondition {
	s.Where(parameterKey+" >= ?", parameterValue)
	return s
}

func (s *SqlCondition) Lt(parameterKey string, parameterValue int) *SqlCondition {
	s.Where(parameterKey+" < ?", parameterValue)
	return s
}

func (s *SqlCondition) Lte(parameterKey string, parameterValue int) *SqlCondition {
	s.Where(parameterKey+" <= ?", parameterValue)
	return s
}

func (s *SqlCondition) Like(parameterKey string, parameterValue string) *SqlCondition {
	s.Where(parameterKey+" LIKE ?", "%"+parameterValue+"%")
	return s
}

func (s *SqlCondition) StartWith(parameterKey string, parameterValue string) *SqlCondition {
	s.Where(parameterKey+" LIKE ?", parameterValue+"%")
	return s
}

func (s *SqlCondition) EndWith(parameterKey string, parameterValue string) *SqlCondition {
	s.Where(parameterKey+" LIKE ?", "%"+parameterValue)
	return s
}

func (s *SqlCondition) In(parameterKey string, parameterValue string) *SqlCondition {
	s.Where(parameterKey+" IN ?", parameterValue)
	return s
}

//func (s *SqlCondition) Ascending(parameterKey string) *SqlCondition {
//	s.OrderBy = append(s.OrderBy, OrderBy{
//		Column:    parameterKey,
//		Ascending: true,
//	})
//	return s
//}
//
//func (s *SqlCondition) Descending(parameterKey string) *SqlCondition {
//	s.OrderBy = append(s.OrderBy, OrderBy{
//		Column:    parameterKey,
//		Ascending: false,
//	})
//	return s
//}

//func (s *SqlCondition) Paginate(page int, pageSize int) *SqlCondition {
//	if s.Paging == nil {
//		s.Paging = &dto.PagingDTO{
//			Page:     page,
//			PageSize: pageSize,
//		}
//	} else {
//		s.Paging.Page = page
//		s.Paging.PageSize = pageSize
//	}
//	return s
//}

func (s *SqlCondition) Build(db *gorm.DB) *gorm.DB {
	//处理顺序：select → where → order → limit → offset
	//select
	if len(s.SelectedColumns) > 0 {
		db = db.Select(s.SelectedColumns)
	}
	//where
	if len(s.ParamPairs) > 0 {
		for _, parameterPair := range s.ParamPairs {
			db = db.Where(parameterPair.ParamKey, parameterPair.ParamValue)
		}
	}
	//order
	var order string
	if s.OrderBy.Ascending == true {
		order = ""
	} else {
		order = " desc"
	}
	column := s.OrderBy.Column
	if column != "" {
		db = db.Order(column + order)
	}
	//limit
	//if s.Paging != nil && s.Paging.PageSize > 0 {
	//	db = db.Limit(s.Paging.PageSize)
	//}

	//offset
	//offset := s.Paging.Offset()
	//if s.Paging != nil && offset > 0 {
	//	db = db.Offset(offset)
	//}

	return db
}

func (s *SqlCondition) Find(db *gorm.DB, model any, output any) error {
	err := s.Build(db).Model(model).Find(&output).Error
	return err
}

// Count 第二个参数应为model类型的指针，如：&model.User{}
// 不理解的话可以看该方法的源码，因为使用了gorm的db.model()方法
func (s *SqlCondition) Count(db *gorm.DB, model any) int {
	result := db.Model(model)
	// where
	if len(s.ParamPairs) > 0 {
		for _, parameterPair := range s.ParamPairs {
			result = result.Where(parameterPair.ParamKey, parameterPair.ParamValue)
		}
	}
	var count int64
	if err := result.Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}
