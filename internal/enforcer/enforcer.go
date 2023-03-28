package enforcer

import (
	"dang_go/internal/database"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strings"
)

// @function: Casbin
// @description: 持久化到数据库  引入自定义规则
// @return: *casbin.Enforcer
func EnforcerTool() *casbin.Enforcer {
	adapter, _ := gormadapter.NewAdapterByDB(database.DB)

	Enforce, err := casbin.NewEnforcer("rbac_model.conf", adapter)

	if err != nil {
		fmt.Printf("%v \n ", "策略初始化失败")
		return Enforce
	}
	// 开启权限认证日志
	//Enforce.EnableLog(true)
	//从DB加载策略
	Enforce.LoadPolicy()

	//	Enforce.AddFunction("ParamsMatch", ParamsMatchFunc)
	//_ = enforcer.LoadPolicy()
	return Enforce
}

// @function: ParamsMatch
// @description: 自定义规则函数
// @param: fullNameKey1 string, key2 string
// @return: bool
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// @function: ParamsMatchFunc
// @description: 自定义规则函数
// @param: args ...interface{}
// @return: interface{}, error
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
