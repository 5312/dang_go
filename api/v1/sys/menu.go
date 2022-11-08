package sys

import (
	"com.example.dang/controller"
	"github.com/kataras/iris/v12/core/router"
)

// 路由

func RegisterRoute(app router.Party) {
	// system 组
	api := app.Party("/sys")
	{
		api.Get("/menus", controller.Menu)
	}
}
