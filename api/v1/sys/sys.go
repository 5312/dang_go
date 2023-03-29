package sys

import (
	"dang_go/controller"
	"github.com/kataras/iris/v12/core/router"
)

// 路由

func RegisterMenuRoute(v1 router.Party) {
	// system 组
	api := v1.Party("/sys")

	{
		// 菜单
		api.Post("/addmenus", controller.InsertMenu)
		api.Delete("/menus/{ID:uint}", controller.DeleteMenu)
		api.Put("/menus/{ID:uint}", controller.UpMenu)
		api.Get("/menus", controller.GetListMenu)
		// 用户
		api.Post("/user/add", controller.InsterUser)
		api.Get("/user/list", controller.GetListUser)
	}

}
