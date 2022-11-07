package sys

import (
	"com.example.dang/controller"
	"github.com/kataras/iris/v12"
)

// 路由

func RegisterRoute(app *iris.Application) {
	// 注册中间件

	// app.Get("/", controller.Menu)

	v1 := app.Party("/v1")

	// 简单分组: v1.
	api := v1.Party("/sys")
	{
		api.Get("/", controller.Menu)
	}
}
