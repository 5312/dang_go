// Package shop: 商品
package shop

import (
	"dang_go/internal/model/shop"
	"dang_go/middleware"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

/*AddLeaseCommodity
* @Description: 添加商品
* @param ctx
 */
func AddLeaseCommodity(ctx iris.Context) {
	// 接收参数
	var data shop.Shop
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	data.FromShops = userInfo.ID
	result, err := data.AddLeaseShop()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}

/*GetLeaseCommodity
* @Description: 我的 租赁商品列表
* @param ctx
* @return {}
 */
func GetLeaseCommodity(ctx iris.Context) {
	//var userInfo middleware.CustomClaims
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	var list shop.Shop
	rest, err := list.GetMyShopList(userInfo)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, rest, "请求成功")
}

/*GetAllCategoryList
* @Description: 获取全部分类商品列表
* @param ctx
 */
func GetAllCategoryList(ctx iris.Context) {
	//var userInfo middleware.CustomClaims
	//userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	id := ctx.URLParam("tid")
	var list shop.Shop
	rest, err := list.GetCategoryShopList(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, rest, "请求成功")
}
