package sys

import (
	"dang_go/controller"
	"dang_go/controller/promoteShop"
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
	// 推广员
	promoter := api.Party("/personnel")
	{
		promoter.Post("/add", promoteShop.AddToPromotePersonnel)            // 添加推广员
		promoter.Get("/list", promoteShop.GetToPromotePersonnel)            // 推广员列表
		promoter.Delete("/{ID:uint}", promoteShop.DeleteToPromotePersonnel) // 删除推广员
		promoter.Put("/{ID:uint}", promoteShop.UpToPromotePersonnel)        // 修改推广员

	}
	// 订单
	{
		promoter.Get("/order/list", promoteShop.OrderList) // 订单列表
	}
}
