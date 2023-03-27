// Package sys 会员
package sys

import (
	"dang_go/controller"
	"github.com/kataras/iris/v12/core/router"
)

// RegisterMemberRoute 会员
func RegisterMemberRoute(app router.Party) {
	// system 组
	api := app.Party("/member")
	{
		// 会员
		api.Get("/list", controller.GetPage)
	}
}
