package api

import (
	"dang_go/middleware"
	"fmt"
	"github.com/spf13/viper"

	"github.com/kataras/iris/v12"
)

// InitUser 初始化api接口
func InitUser() {
	/* 1. 初始化iris */
	app := iris.New()
	/* 2. 中间件 */
	middleware.InitMiddleware(app)

	// 注册系统路由
	InitSysRouter(app)

	// 读取conf
	str := ":"
	port := viper.Get("Conf.Port")
	res := fmt.Sprintf("%s%s", str, port)
	_ = app.Run(iris.Addr(res))
}
