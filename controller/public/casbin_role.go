package public

import (
	"dang_go/internal/model"
	"dang_go/tools/app"
	"fmt"
	"github.com/kataras/iris/v12"
)

func Addrole(ctx iris.Context) {

	// 接收参数
	var data model.CasbinModel
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "请输入参数")
		return
	}

	err := data.Create()
	fmt.Printf("%v \n", data)
	if err != nil {
		app.Error(ctx, -1, err, "登录失败")
		return
	}

	app.OK(ctx, 1, "登录成功")
}
