package util

import (
	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"learn-go/model"
)

func NewEnforcer() *casbin.Enforcer {
	adapter, _ := gormAdapter.NewAdapterByDB(model.DB)
	enforcer, _ := casbin.NewEnforcer("./config/casbin_model.conf", adapter)
	_ = enforcer.LoadPolicy()
	return enforcer
}
