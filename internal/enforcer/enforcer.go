package enforcer

import (
	"dang_go/internal/database"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strings"
)

/*EnforcerTool
* @Description: 持久化到数据库  引入自定义规则
* @return *casbin.Enforcer
 */
//goland:noinspection GoNameStartsWithPackageName
func EnforcerTool() *casbin.Enforcer {
	adapter, _ := gormadapter.NewAdapterByDB(database.DB)

	Enforce, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)

	if err != nil {
		fmt.Printf("%v \n ", "策略初始化失败")
		return Enforce
	}
	// 开启权限认证日志
	//Enforce.EnableLog(true)
	//从DB加载策略
	errs := Enforce.LoadPolicy()
	if errs != nil {
		return nil
	}

	//	Enforce.AddFunction("ParamsMatch", ParamsMatchFunc)
	//_ = enforcer.LoadPolicy()
	return Enforce
}

/*ParamsMatch
* @Description: 自定义规则函数
* @param fullNameKey1  string, key2 string
* @param key2
* @return bool
 */
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

/*ParamsMatchFunc
* @Description: 自定义规则函数
* @param args
* @return interface{}
* @return error
 */
//goland:noinspection GoUnusedExportedFunction
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
