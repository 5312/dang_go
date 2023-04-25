package apialipay

import (
	"dang_go/controller"
	"dang_go/controller/conalipay"
	"dang_go/pay"
	"github.com/kataras/iris/v12/core/router"
)

func RegisterAppLetsRoute(v1 router.Party) {
	api := v1.Party("/applets")
	//banner
	{
		// banner
		api.Post("/home-img/add", conalipay.AddBanner)         // 添加banner
		api.Get("/home-img/list", conalipay.GetBanner)         // 获取banner
		api.Put("/home-img/{ID:uint}", conalipay.UpdateBanner) // 修改banner

		// 商品添加至 推荐
		api.Get("/recommend/add", conalipay.AddRecommend)
		api.Get("/recommend/list", conalipay.GetRecommend) // 获取活动上坪

	}

	// 支付宝sdk
	{
		api.Post("/authentication", conalipay.AlipayUserCertifyOpenInitializeRequest) // 人脸身份认证
		api.Post("/save-idCard", controller.SaveIdCard)                               // 保存身份信息
		api.Post("/upLngLat", controller.SaveUpLngLat)                                // 保存下单时经纬度

		api.Post("/pay-order/{OID:uint}/{RID:uint}", pay.Init) // 支付 oid订单id 和账单rid
		api.Post("/pay/trade", pay.Trade)                      // 统一收单交易创建
		api.Post("/pay/freeze", pay.Freeze)                    // 线上资金授权冻结接口

		api.Post("/save-order", pay.SaveOrder)                // 保存订单
		api.Get("/order/list", pay.OrderPage)                 // 查看订单列表
		api.Get("/order/{OID:uint}/detail", pay.OrderDetail)  // 查看订单列表
		api.Post("/order/{OID:uint}/cancel", pay.OrderCancel) // 用户取消订单
		api.Post("/order/sync", pay.OrderSync)                // 订单数据同步

		api.Post("/decryptPhoneNum", conalipay.DecryptPhoneNum) // 手机号解密

	}
}
