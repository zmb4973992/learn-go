package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"learn-go/dao"
)

func Init() {
	adapter, _ := gormadapter.NewAdapterByDB(dao.DB)
	enforcer, _ := casbin.NewEnforcer("./config/casbin_model.conf", adapter)
	err := enforcer.LoadPolicy()
	if err != nil {
		panic("casbin加载策略失败，请重试")
	}
}
