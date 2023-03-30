package controller

import (
	"dang_go/internal/model/shop"
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

type LoginParams struct {
	Name     string
	Password string
}

/*Login
* @Description: 平台登录
* @param ctx
 */
func Login(ctx iris.Context) {
	var user system.User
	// 接收参数
	var data LoginParams
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "请输入参数")
		return
	}

	name := data.Name
	password := data.Password
	success, err := user.Login(name, password)
	if err != nil {
		app.Error(ctx, -1, err, "登录失败")
		return
	}

	app.OK(ctx, success, "登录成功")
}

/*ShopLogin
* @Description: 商户登录
* @param ctx
 */
func ShopLogin(ctx iris.Context) {
	var user shop.Merchant
	// 接收参数
	var data LoginParams
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "请输入参数")
		return
	}

	name := data.Name
	password := data.Password
	success, err := user.Login(name, password)
	if err != nil {
		app.Error(ctx, -1, err, "登录失败")
		return
	}

	app.OK(ctx, success, "登录成功")
}
