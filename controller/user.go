package controller

import (
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"fmt"
	"github.com/kataras/iris/v12"
)

// InsterUser
//TODO: 账号 account 不能重复添加验证
//TODO: 密码加密 使用bcrypt

func InsterUser(ctx iris.Context) {
	// 接收参数
	var data system.User
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	fmt.Println(data)

	result, err := data.Create()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}

/*GetListUser
* @Description 获取用户信息
* @Tags 用户
* @Accept  json
* @Produce  json
* @Param id path int true "用户ID"
 */
func GetListUser(ctx iris.Context) {
	// Get all records
	var data system.User
	name := ctx.URLParam("name")

	result, err := data.GetPage(name)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, result, "查询成功")
}
