package conalipay

import (
	"dang_go/internal/model/shop"
	"dang_go/tools/app"
	"github.com/gogf/gf/util/gconv"
	"github.com/kataras/iris/v12"
)

/*AddCategory
* @Description: 添加分类
* @param ctx
 */
func AddCategory(ctx iris.Context) {
	var data shop.Category
	if err := ctx.ReadJSON(&data); err != nil {

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

/*Update
* @Description: 修改信息
* @param ctx
 */
func Update(ctx iris.Context) {
	// 接收参数
	var update shop.Category
	id, _ := ctx.Params().GetUint("ID")
	if err := ctx.ReadJSON(&update); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	// 接受参数
	result, err := update.Update(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "修改成功")
}

/*Delete 删 */
func Delete(ctx iris.Context) {
	var data shop.Category
	id, _ := ctx.Params().GetUint("ID")

	_, err := data.Delete(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, id, "删除成功")
}

/*GetTreeList
* @Description: 获取树状数据
* @param ctx
 */
func GetTreeList(ctx iris.Context) {
	// Get all records
	var data shop.Category
	name := ctx.URLParam("name")
	startime := ctx.URLParam("startTime")
	endtime := ctx.URLParam("startTime")

	result, err := data.GetPage(name, startime, endtime)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	arr := TreeNode(result, 0)
	app.OK(ctx, arr, "查询成功")

}

/*TreeNode 转换树结构 */
func TreeNode(list []shop.Category, pid uint) []shop.Categorys {
	var array []shop.Categorys

	for _, item := range list {
		if item.ParentId == pid {
			var children shop.Categorys

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
