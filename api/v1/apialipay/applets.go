package apialipay

import (
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
		api.Post("/pay-order/{OID:uint}/{RID:uint}", pay.Init)                        // 支付 oid订单id 和账单rid
		api.Post("/save-order", pay.SaveOrder)                                        // 保存订单

	}
}
