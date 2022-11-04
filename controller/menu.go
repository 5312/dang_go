package controller

import "github.com/kataras/iris/v12"

func Menu(ctx iris.Context) {
	ctx.View("index.html")
}

/* 默认显示页面 */
func Index(ctx iris.Context) {
	ctx.View("index.html")
}
