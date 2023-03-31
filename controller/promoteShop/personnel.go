package promoteShop

import (
	"dang_go/internal/model/promoter"
	"dang_go/middleware"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

/*AddToPromotePersonnel
* @Description: 添加推广员
* @param ctx
 */
func AddToPromotePersonnel(ctx iris.Context) {
	// 接收参数
	var data promoter.Personnel
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	//fmt.Printf("%v", data)
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	data.FromPromoter = userInfo.ID

	result, err := data.Create()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}

/*DeleteToPromotePersonnel
* @Description: 删除推广员
* @param ctx
* @return {}
 */
func DeleteToPromotePersonnel(ctx iris.Context) {
	var data promoter.Personnel
	id, _ := ctx.Params().GetUint("ID")

	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	data.FromPromoter = userInfo.ID

	err := data.Delete(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, id, "删除成功")
}

/*GetToPromotePersonnel
* @Description: 推广员列表
* @param ctx
 */
func GetToPromotePersonnel(ctx iris.Context) {
	//var userInfo middleware.CustomClaims
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	var list promoter.Personnel
	rest, err := list.GetMyPersonnelList(userInfo)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, rest, "请求成功")
}

/*UpToPromotePersonnel
* @Description: 修改推广员信息
* @param ctx
 */
func UpToPromotePersonnel(ctx iris.Context) {
	// 接收参数
	var data promoter.Personnel
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
