package promoteShop

import (
	"dang_go/internal/model/shop"
	"dang_go/middleware"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

/*OrderList
* @Description: 订单列表
* @param ctx
 */
func OrderList(ctx iris.Context) {
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)

	var order shop.Order
	result, err := order.GetMerchantOrder(userInfo.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "添加成功")

}
