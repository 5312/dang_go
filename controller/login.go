package controller

import (
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	var user system.User
	name := ctx.URLParam("name")
	password := ctx.URLParam("password")

	success, err := user.Login(name, password)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, success, "删除成功")
}
