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
		api.Post("/merchant/add", controller.InsterShop)      // 开户
		api.Get("/merchant/list", controller.GetListShop)     // 商家列表
		api.Put("/merchant/{ID:uint}", controller.UpMerchant) // 商家列表

	}
	// 商品管理
	{
		// 租赁地址
		api.Post("/address/lease/add/{ID:uint}", controller.AddLeaseAddress)
	}
}
