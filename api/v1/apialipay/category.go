// Package apialipay
// @Description:
package apialipay

import (
	"dang_go/controller/conalipay"
	"github.com/kataras/iris/v12/core/router"
)

func RegisterShopCategoryRoute(v1 router.Party) {
	api := v1.Party("/alipay")
	//分类
	{
		// 分类
		api.Post("/category/add", conalipay.AddCategory)
		api.Get("/category/list", conalipay.GetTreeList)
		api.Put("/category/{ID:uint}", conalipay.Update)
		api.Delete("/category/{ID:uint}", conalipay.Delete)

	}
}
