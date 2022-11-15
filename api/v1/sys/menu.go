package sys

import (
	. "dang_go/controller"

	"github.com/kataras/iris/v12/core/router"
)

// 路由

func RegisterRoute(app router.Party) {
	// system 组
	api := app.Party("/sys")
	{

		menu := &Menus{
			Name: "菜单",
		}

		api.Post("/addmenus", menu.AddMenu)
		api.Delete("/menus/{ID:int}", menu.DeleteMenu)
		api.Get("/menus", menu.GetListMenu)
	}
}
