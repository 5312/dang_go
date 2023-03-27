package controller

import (
	"dang_go/internal/model/shop"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

/*
添加商户
*/
func InsterShop(ctx iris.Context) {
	// 接收参数
	var data shop.Merchant
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	result, err := data.AddShop()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}

func GetListShop(ctx iris.Context) {
	// Get all records
	var data shop.Merchant

	result, err := data.GetPage()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, result, "查询成功")
}
