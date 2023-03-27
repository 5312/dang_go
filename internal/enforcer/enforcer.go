package enforcer

import (
	"dang_go/internal/database"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func EnforcerTool() *casbin.Enforcer {
	fmt.Printf("%v \n------Casbin", "-----")
	adapter, _ := gormadapter.NewAdapterByDB(database.DB)
	Enforce, err := casbin.NewEnforcer("rbac_model.conf", adapter)

	// 开启权限认证日志
	Enforce.EnableLog(true)

	if err != nil {
		// 加载数据库中的策略
		loaderr := Enforce.LoadPolicy()
		if loaderr != nil {
			panic(loaderr)
		}
		return Enforce
	}
	fmt.Printf("v \n", err)
	return nil
}
