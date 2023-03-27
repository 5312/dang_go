package controller

import (
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

func GetPage(ctx iris.Context) {
	// Get all records
	var data system.Member

	result, err := data.GetPage()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, result, "查询成功")
}
