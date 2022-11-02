package router

import (
	"github.com/kataras/iris/v12"
)

// 路由

func RegisterRoute(app *iris.Application) {
	// 注册中间件

	v1 := app.Party("/v1")

	// 简单分组: v1.
	/* api := v1.Party("/api")
	{
		api.Get("/", controller.Users)
	}

	admin := v1.Party("/admin")
	{
		admin.Get("/", controller.Users)
	} */
}
