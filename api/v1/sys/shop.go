package sys

import (
	"dang_go/controller/shop"
	"github.com/kataras/iris/v12/core/router"
)

// 路由

func RegisterShopRoute(v1 router.Party) {
	// system 组
	api := v1.Party("/shop")
	//商户
	{
		// 商户
		api.Post("/merchant/add", shop.InsterShop)      // 开户
		api.Get("/merchant/list", shop.GetListShop)     // 商家列表
		api.Put("/merchant/{ID:uint}", shop.UpMerchant) // 商家列表

	}
	// 商品管理
	{
		// 发布 租赁 商品
		api.Post("/rent/add", shop.AddLeaseCommodity)
		api.Get("/rent/list", shop.GetLeaseCommodity)
		api.Get("/rent/{ID:uint}", shop.GetShopDetail)

		// 租赁地址
		api.Post("/address/lease/add/{ID:uint}", shop.AddLeaseAddress)
		// 归还地址
		api.Post("/address/return/add/{ID:uint}", shop.AddReturnAddress)

	}
}
