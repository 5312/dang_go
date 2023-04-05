package shop

import (
	"dang_go/internal/model/shop"
	"dang_go/middleware"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

// 查询自己分配的订单列表

/*GetAllOrderPage
* @Description: 获取全部订单
* @param ctx
 */
func GetAllOrderPage(ctx iris.Context) {
	var statusType = ctx.URLParam("status")
	// 接收参数
	var data shop.Order
	result, err := data.GetAllPage(statusType)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}

/*GetMerchantOrder
* @Description: 商家累计订单数
* @param ctx
* @return {}
 */
func GetMerchantOrderCount(ctx iris.Context) {
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)

	var order shop.Order
	result, err := order.GetCount(userInfo.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}

/*GetMerchantOrder
* @Description: 商家查询自订单列表
* @param ctx
 */
func GetMerchantOrder(ctx iris.Context) {
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)

	var order shop.Order
	result, err := order.GetMerchantOrder(userInfo.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")
}
