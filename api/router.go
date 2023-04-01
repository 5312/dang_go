package api

import (
	"dang_go/api/v1/apialipay"
	"dang_go/api/v1/public"
	"dang_go/api/v1/sys"
	"dang_go/controller"
	_ "dang_go/docs"
	"github.com/kataras/iris/v12"
)

// InitSysRouter 初始化api接口
func InitSysRouter(app *iris.Application) {
	// V1
	v1 := app.Party("/v1")
	{
		// 登录  跳过jwt
		v1.Post("/login", controller.Login)                  // 平台
		v1.Post("/alipay/login", controller.AlipayLogin)     // 支付宝
		v1.Post("/shop/login", controller.ShopLogin)         // 商户
		v1.Post("/promoter/login", controller.PromoterLogin) // 推广商

	}
	{
		// sys
		sys.RegisterMenuRoute(v1)     // 菜单
		sys.RegisterShopRoute(v1)     // 商户
		sys.RegisterMemberRoute(v1)   // 会员
		sys.RegisterPromoterRoute(v1) // 推广商

		// 小程序
		apialipay.RegisterShopCategoryRoute(v1) // 分类
		apialipay.RegisterAppLetsRoute(v1)      // banner图
		//公共接口
		public.RegisterPublicRoute(v1)
	}

}
