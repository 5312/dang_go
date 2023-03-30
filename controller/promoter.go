package controller

import (
	"dang_go/internal/model/promoter"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

// AddPromoter 添加
func AddPromoter(ctx iris.Context) {
	var data promoter.Promoter
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

/*GetPromoterPageList
* @Description: 查询
* @param ctx
 */
func GetPromoterPageList(ctx iris.Context) {

	var data promoter.Promoter
	result, err := data.GetPage("")
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "")

}

/*DeleteFormId
* @Description: 删除
* @param ctx
 */
func DeleteFormId(ctx iris.Context) {
	var data promoter.Promoter
	id, _ := ctx.Params().GetUint("ID")
	err := data.Delete(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, id, "删除成功")
}

/*PutData
* @Description: 修改
* @param ctx
 */
func PutData(ctx iris.Context) {
	var data promoter.Promoter
	// 读取id
	id, _ := ctx.Params().GetUint("ID")
	// 读取参数
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
