package public

import (
	"dang_go/controller/public"
	"github.com/kataras/iris/v12/core/router"
)

// 公共路由 v1
func RegisterPublicRoute(app router.Party) {
	// system 组
	api := app.Party("/public")
	//商户
	{
		// 上传图片
		api.Post("/upload/image", public.UploadImage)
		api.Post("/permissions/add", public.Addrole)

	}
}
