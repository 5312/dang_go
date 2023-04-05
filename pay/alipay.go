package pay

import (
	"dang_go/internal/model/shop"
	"dang_go/middleware"
	"dang_go/tools/app"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/kataras/iris/v12"
)

/*Init
* @Description: 订单支付
* @param ctx
 */
func Init(ctx iris.Context) {
	xlog.Info("GoPay Version: ", gopay.Version)

	oid, _ := ctx.Params().GetUint("OID")
	rid, _ := ctx.Params().GetUint("RID")

	fmt.Printf("oid%v \n", oid)
	fmt.Printf("rid:%v \n", rid)
}

/*SaveOrder
* @Description: 保存订单信息
* @param ctx
 */
func SaveOrder(ctx iris.Context) {
	var params shop.Order
	if err := ctx.ReadJSON(&params); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	id := ctx.Values().Get("claims").(*middleware.CustomClaims)

	ormCr, err := params.Create(id.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, ormCr, "添加成功")
}

/*OrderPage
* @Description: 订单列表
 */
func OrderPage(ctx iris.Context) {
	var order shop.Order
	// 会员id
	id := ctx.Values().Get("claims").(*middleware.CustomClaims)

	result, err := order.GetPage(id.ID)

	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "查询成功")
}

/*OrderDetail
* @Description: 查询订单详情
* @param ctx
 */
func OrderDetail(ctx iris.Context) {
	var order shop.Order
	// 会员id
	id, err := ctx.Params().GetUint("OID")

	result, errs := order.GetOrderDetail(id)

	if errs != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "查询成功")
}

/*OrderCancel
* @Description: 取消订单
* @param ctx
* @return {}
 */
func OrderCancel(ctx iris.Context) {
	var order shop.Order
	// 订单id
	id, err := ctx.Params().GetUint("OID")

	result, errs := order.Cancel(id)

	if errs != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "取消成功")
}
