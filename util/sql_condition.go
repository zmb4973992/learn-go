package util

import (
	"gorm.io/gorm"
	"learn-go/dto"
)

type SqlCondition struct {
	SelectedColumns []string
	ParamPairs      []ParamPair
	OrderBy         dto.OrderByDTO
	Paging          dto.PagingDTO
}

type ParamPair struct {
	ParamKey   string //查询参数的名称，如 age>=, name include, id=
	ParamValue any    //查询参数的值
}

// NewSqlCondition 生成自定义的查询条件,参数可不填
//必须为指针，因为下面的方法要用到指针进行修改入参
func NewSqlCondition() *SqlCondition {
	return &SqlCondition{
		Paging: dto.PagingDTO{
			Page:     1,
			PageSize: 20,
		},
	}
}

// Where 给SqlCondition自定义where方法，将参数保存到ParameterPair中
func (s *SqlCondition) Where(key string, value any) *SqlCondition {
	s.ParamPairs = append(s.ParamPairs, ParamPair{
		ParamKey:   key,
		ParamValue: value,
	})
	return s
}

func (s *SqlCondition) Equal(paramKey string, paramValue any) *SqlCondition {
	s.Where(paramKey+" = ? ", paramValue)
	return s
}

func (s *SqlCondition) NotEqual(paramKey string, paramValue any) *SqlCondition {
	s.Where(paramKey+" <> ?", paramValue)
	return s
}

func (s *SqlCondition) Gt(paramKey string, paramValue int) *SqlCondition {
	s.Where(paramKey+" > ?", paramValue)
	return s
}

func (s *SqlCondition) Gte(paramKey string, paramValue int) *SqlCondition {
	s.Where(paramKey+" >= ?", paramValue)
	return s
}

func (s *SqlCondition) Lt(paramKey string, paramValue int) *SqlCondition {
	s.Where(paramKey+" < ?", paramValue)
	return s
}

func (s *SqlCondition) Lte(paramKey string, paramValue int) *SqlCondition {
	s.Where(paramKey+" <= ?", paramValue)
	return s
}

// Include 和Like为相同方法
func (s *SqlCondition) Include(paramKey string, paramValue string) *SqlCondition {
	s.Where(paramKey+" LIKE ?", "%"+paramValue+"%")
	return s
}

// Like 和Include为相同方法
func (s *SqlCondition) Like(paramKey string, paramValue string) *SqlCondition {
	s.Where(paramKey+" LIKE ?", "%"+paramValue+"%")
	return s
}

func (s *SqlCondition) StartWith(paramKey string, paramValue string) *SqlCondition {
	s.Where(paramKey+" LIKE ?", paramValue+"%")
	return s
}

func (s *SqlCondition) EndWith(paramKey string, paramValue string) *SqlCondition {
	s.Where(paramKey+" LIKE ?", "%"+paramValue)
	return s
}

func (s *SqlCondition) In(paramKey string, paramValue string) *SqlCondition {
	s.Where(paramKey+" IN ?", paramValue)
	return s
}

//func (s *SqlCondition) Ascending(parameterKey string) *SqlCondition {
//	s.OrderByDTO = append(s.OrderByDTO, OrderByDTO{
//		OrderByColumn:    parameterKey,
//		Ascending: true,
//	})
//	return s
//}
//
//func (s *SqlCondition) Descending(parameterKey string) *SqlCondition {
//	s.OrderByDTO = append(s.OrderByDTO, OrderByDTO{
//		OrderByColumn:    parameterKey,
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
	if s.OrderBy.OrderByColumn != "" {
		if s.OrderBy.Desc == true {
			db = db.Order(s.OrderBy.OrderByColumn + " desc")
		}
		db = db.Order(s.OrderBy.OrderByColumn)
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
	err := s.Build(db).Debug().Model(model).Find(&output).Error
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
