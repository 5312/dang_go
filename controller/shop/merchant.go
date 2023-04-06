package shop

import (
	"dang_go/internal/model/shop"
	"dang_go/middleware"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

/*InsterShop
* @Description: 添加商户
* @param ctx
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

/*GetListShop
* @Description: 获取商户列表
* @param ctx
 */
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

/*UpMerchant
* @Description: 商户修改
* @param ctx
 */
func UpMerchant(ctx iris.Context) {
	// 接收参数
	var data shop.Merchant
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

/*AddLeaseAddress
* @Description: 商户 添加 修改 租赁地址
* @param ctx
 */
func AddLeaseAddress(ctx iris.Context) {
	// 接收参数
	var data shop.Medium
	id, _ := ctx.Params().GetUint("ID")
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	// 将用户添加到切片中
	update, err := data.UpdateAddress(id, data, "address_lease")

	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, update, "操作成功")
}

// AddReturnAddress
// @Description: {商户 添加 修改 归还地址}
// @param ctx
func AddReturnAddress(ctx iris.Context) {
	// 接收参数
	var data shop.Medium
	id, _ := ctx.Params().GetUint("ID")
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	update, err := data.UpdateAddress(id, data, "address_return")

	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, update, "操作成功")
}

/*AddMerchantMan
* @Description: 添加商家成员
* @param ctx
* @return {}
 */
func AddMerchantMan(ctx iris.Context) {
	var params shop.MerchantMan
	if err := ctx.ReadJSON(&params); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	res, err := params.Create(userInfo.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, res, "添加成功")
}

/*GetMerchantMan
* @Description: 成员列表
* @param ctx
 */
func GetMerchantMan(ctx iris.Context) {
	var params shop.MerchantMan

	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	res, err := params.GetList(userInfo.ID)

	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, res, "添加成功")
}
