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
		api.Post("/merchant/add", shop.InsterShop)                      // 开户
		api.Get("/merchant/list", shop.GetListShop)                     // 商家列表
		api.Put("/merchant/{ID:uint}", shop.UpMerchant)                 // 商家列表
		api.Post("/merchant/{SID:uint}/detail", shop.GetMerchantDetail) // 查询商家
		api.Get("/merchant/{SID:uint}/shop", shop.GetMerchantShop)      // 查询商家商品
		api.Get("/merchant/order", shop.GetMerchantOrderCount)          // 查询商家 累计订单数
		api.Get("/merchant/order", shop.GetMerchantOrder)               // 查询商家订单列表

	}
	// 商品管理
	{
		// 发布 租赁 商品
		api.Post("/rent/add", shop.AddLeaseCommodity)  // 添加商品
		api.Get("/rent/list", shop.GetLeaseCommodity)  // 商品列表
		api.Get("/rent/{ID:uint}", shop.GetShopDetail) // 商品详情
		api.Get("/order/all", shop.GetAllOrderPage)    // 商品详情

		// 租赁地址
		api.Post("/address/lease/add/{ID:uint}", shop.AddLeaseAddress)
		// 归还地址
		api.Post("/address/return/add/{ID:uint}", shop.AddReturnAddress)

	}
}
