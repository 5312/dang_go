package sys

import (
	"dang_go/controller"
	"github.com/kataras/iris/v12/core/router"
)

// 路由

func RegisterShopRoute(v1 router.Party) {
	// system 组
	api := v1.Party("/shop")
	//商户
	{
		// 商户
		api.Post("/merchant/add", controller.InsterShop)
		api.Get("/merchant/list", controller.GetListShop)

	}
}
