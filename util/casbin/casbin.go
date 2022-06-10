package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"learn-go/util"
)

func Init() {
	a, _ := gormadapter.NewAdapterByDB(util.DB)
	e, _ := casbin.NewEnforcer("./config/rbac_model.conf", a)
	e.LoadPolicy()
	res, _ := e.Enforce("a", "b", "c")
	fmt.Println(res)
}
