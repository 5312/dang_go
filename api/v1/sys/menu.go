package sys

import (
	"dang_go/controller"
	"github.com/kataras/iris/v12/core/router"
)

// 路由

func RegisterRoute(app router.Party) {
	// system 组
	api := app.Party("/sys")
	{

		api.Post("/addmenus", controller.InsertMenu)
		api.Delete("/menus/{ID:uint}", controller.DeleteMenu)
		api.Put("/menus/{ID:uint}", controller.UpMenu)
		api.Get("/menus", controller.GetListMenu)
	}
}
