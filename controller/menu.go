package controller

import (
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/kataras/iris/v12"
)

/*InsertMenu 增 */
func InsertMenu(ctx iris.Context) {
	// 接收参数
	var data system.Menu
	if err := ctx.ReadJSON(&data); err != nil {
		//panic(err.Error())
		app.Error(ctx, -1, err, "")
		return
	}
	result, err := data.Create()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}

/*DeleteMenu 删 */
func DeleteMenu(ctx iris.Context) {
	var data system.Menu
	id, _ := ctx.Params().GetUint("ID")

	_, err := data.Delete(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, id, "删除成功")
}

/*UpMenu 改 */
func UpMenu(ctx iris.Context) {
	// 接收参数
	var data system.Menu
	id, _ := ctx.Params().GetUint("ID")
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	result, err := data.Update(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "修改成功")
}

/*GetListMenu 查 */
// name:
//startTime:
//endTime:
func GetListMenu(ctx iris.Context) {
	// Get all records
	var data system.Menu
	name := ctx.URLParam("name")
	startime := ctx.URLParam("startTime")
	endtime := ctx.URLParam("startTime")

	fmt.Println(name)
	fmt.Println(startime)
	fmt.Println(endtime)

	result, err := data.GetPage(name, startime, endtime)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	arr := TreeNode(result, 0)
	app.OK(ctx, arr, "查询成功")
}

/*TreeNode 转换树结构 */
func TreeNode(list []system.Menu, pid uint) []system.Menus {
	var array []system.Menus

	for _, item := range list {
		if item.ParentId == pid {
			var children system.Menus

			err := gconv.Struct(item, &children)
			if err != nil {
				panic(err)
			}
			children.Children = TreeNode(list, item.Model.ID)
			array = append(array, children)
		}
	}
	return array
}
