package public

import (
	"dang_go/controller/public"
	"github.com/kataras/iris/v12/core/router"
)

// 公共路由
func RegisterShopRoute(app router.Party) {
	// system 组
	api := app.Party("/public")
	//商户
	{
		// 商户
		api.Post("/upload/image", public.UploadImage)

	}
}
