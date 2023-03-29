package public

import (
	"dang_go/controller/public"
	"github.com/kataras/iris/v12/core/router"
)

// 公共路由 v1
func RegisterPublicRoute(v1 router.Party) {
	// system 组
	api := v1.Party("/public")
	//商户
	{
		// 上传图片
		api.Post("/upload/image", public.UploadImage)
		api.Post("/permissions/add", public.Addrole)

	}
}
