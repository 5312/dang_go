package api

import (
	"dang_go/api/v1/apialipay"
	"dang_go/api/v1/public"
	"dang_go/api/v1/sys"
	"dang_go/controller"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	_ "dang_go/docs"
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
		//权限验证
		//v1.Use(middleware.JWTAuth)
		// sys
		sys.RegisterMenuRoute(v1)     // 菜单
		sys.RegisterShopRoute(v1)     // 商户
		sys.RegisterMemberRoute(v1)   // 会员
		sys.RegisterPromoterRoute(v1) // 推广商

		// 小程序
		apialipay.RegisterShopCategoryRoute(v1) // 分类
		//公共接口
		public.RegisterPublicRoute(v1)
	}
	/* swagger文档*/
	// 指向swagger init生成文档的路径
	//config := &swagger.Config{
	//	URL:         "http://localhost:87/docs/swagger.json",
	//	DeepLinking: true,
	//}
	//app.Get("/swagger/*any", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
	app.Get("/swagger/index.html", swagger.WrapHandler(swaggerFiles.Handler))
}
