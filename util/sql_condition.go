package util

import (
	"gorm.io/gorm"
)

type SqlCondition struct {
	SelectedColumns []string
	ParameterPairs  []ParameterPair
	OrderByColumns  []OrderBy
	PagingRule      *Paging
}

type ParameterPair struct {
	ParameterKey   string //查询参数的名称，如 age>=, name include, id=
	ParameterValue any    //查询参数的值
}

type OrderBy struct {
	Column    string //排序字段
	Ascending bool   //是否为升序（从小到大）
}

// NewSqlCondition 生成自定义的查询条件,参数可不填
func NewSqlCondition(selectedColumns ...string) *SqlCondition {
	result := &SqlCondition{}
	if len(selectedColumns) > 0 {
		result.SelectedColumns = append(result.SelectedColumns, selectedColumns...)
	}
	return result
}

// Where 给SqlCondition自定义where方法，将参数保存到ParameterPair中
func (s *SqlCondition) Where(parameterKey string, parameterValue any) *SqlCondition {
	s.ParameterPairs = append(s.ParameterPairs, ParameterPair{
		ParameterKey:   parameterKey,
		ParameterValue: parameterValue,
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

func (s *SqlCondition) GreaterThan(parameterKey string, parameterValue int) *SqlCondition {
	s.Where(parameterKey+" > ?", parameterValue)
	return s
}

func (s *SqlCondition) GreaterThanOrEqual(parameterKey string, parameterValue int) *SqlCondition {
	s.Where(parameterKey+" >= ?", parameterValue)
	return s
}

func (s *SqlCondition) LessThan(parameterKey string, parameterValue int) *SqlCondition {
	s.Where(parameterKey+" < ?", parameterValue)
	return s
}

func (s *SqlCondition) LessThanOrEqual(parameterKey string, parameterValue int) *SqlCondition {
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

func (s *SqlCondition) Ascending(parameterKey string) *SqlCondition {
	s.OrderByColumns = append(s.OrderByColumns, OrderBy{
		Column:    parameterKey,
		Ascending: true,
	})
	return s
}

func (s *SqlCondition) Descending(parameterKey string) *SqlCondition {
	s.OrderByColumns = append(s.OrderByColumns, OrderBy{
		Column:    parameterKey,
		Ascending: false,
	})
	return s
}

func (s *SqlCondition) Paginate(page int, pageSize int) *SqlCondition {
	if s.PagingRule == nil {
		s.PagingRule = &Paging{
			Page:     page,
			PageSize: pageSize,
		}
	} else {
		s.PagingRule.Page = page
		s.PagingRule.PageSize = pageSize
	}
	return s
}

func (s *SqlCondition) Build(db *gorm.DB) *gorm.DB {
	//处理顺序：select → where → order → limit → offset
	//选择要传递的字段,select columns
	if len(s.SelectedColumns) > 0 {
		db = db.Select(s.SelectedColumns)
	}
	//where
	if len(s.ParameterPairs) > 0 {
		for _, parameterPair := range s.ParameterPairs {
			db = db.Where(parameterPair.ParameterKey, parameterPair.ParameterValue)
		}
	}
	//order
	if len(s.OrderByColumns) > 0 {
		for _, orderByColumn := range s.OrderByColumns {
			if orderByColumn.Ascending == true {
				db = db.Order(orderByColumn.Column + "ASC")
			} else {
				db = db.Order(orderByColumn.Column + "DESC")
			}
		}
	}
	//limit
	if s.PagingRule != nil && s.PagingRule.PageSize > 0 {
		db = db.Limit(s.PagingRule.PageSize)
	}

	//offset
	offset := s.PagingRule.Offset()
	if s.PagingRule != nil && offset > 0 {
		db = db.Offset(offset)
	}
	return db
}

func (s *SqlCondition) Find(db *gorm.DB) (any, error) {
	var output any
	err := s.Build(db).Find(&output).Error
	return output, err
}

// Count 第二个参数应为model类型的指针，如：&model.User{}
// 不理解的话可以看该方法的源码，因为使用了gorm的db.model()方法
func (s *SqlCondition) Count(db *gorm.DB, model any) int {
	result := db.Model(model)
	// where
	if len(s.ParameterPairs) > 0 {
		for _, parameterPair := range s.ParameterPairs {
			result = result.Where(parameterPair.ParameterKey, parameterPair.ParameterValue)
		}
	}
	var count int64
	if err := result.Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}
