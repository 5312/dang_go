// Package sys 会员
package sys

import (
	"dang_go/controller"
	"github.com/kataras/iris/v12/core/router"
)

// RegisterMemberRoute 会员
func RegisterMemberRoute(v1 router.Party) {
	// system 组
	api := v1.Party("/member")
	{
		// 会员
		api.Get("/list", controller.GetPage)

		// 添加地址
		api.Post("/address/add", controller.Address)
		// 获取地址
		api.Get("/address/list", controller.AddressList)

	}
}
