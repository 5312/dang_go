package sys

import (
	"dang_go/controller"
	"github.com/kataras/iris/v12/core/router"
)

// RegisterPromoterRoute 路由 /v1
func RegisterPromoterRoute(v1 router.Party) {
	// system 组
	api := v1.Party("/promoter")
	//推广商
	{
		// 推广商
		api.Post("/add", controller.AddPromoter)
		api.Get("/list", controller.GetPromoterPageList)
		api.Delete("/{ID:uint}", controller.DeleteFormId)
		api.Put("/{ID:uint}", controller.PutData)

	}
}