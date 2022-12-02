package sys

import (
	"dang_go/controller"
	"github.com/kataras/iris/v12/core/router"
)

// RegisterUserRoute : 路由
func RegisterUserRoute(app router.Party) {
	// system 组
	api := app.Party("/sys")
	{

		api.Post("/user/add", controller.InsterUser)
		//api.Delete("/menus/{ID:uint}", controller.DeleteMenu)
		//api.Put("/menus/{ID:uint}", controller.UpMenu)
		api.Get("/user/list", controller.GetListUser)
	}
}
