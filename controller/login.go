package controller

import (
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

type LoginParams struct {
	Name     string
	Password string
}

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
