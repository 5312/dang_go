package pay

import (
	"dang_go/controller/conalipay"
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

type AuthParams struct {
	OutTradeNo  string `json:"out_trade_no"`
	TotalAmount string `json:"total_amount"`
	Subject     string `json:"subject"`
	BuyerId     string `json:"buyer_id"`
}

type FreeParams struct {
	OutTradeNo     string `json:"out_trade_no"`
	OutRequestNo   string `json:"out_request_no"`
	OrderTitle     string `json:"order_title"`
	Amount         int    `json:"amount"`
	ProductCode    string `json:"product_code"`
	PayeeUserId    string `json:"payee_user_id"`
	TimeoutExpress string `json:"timeout_express"`
}

/*Freeze
* @Description: 生成资金冻结订单
* @param ctx
 */
func Freeze(ctx iris.Context) {

	var userData FreeParams
	if err := ctx.ReadJSON(&userData); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	client, err := conalipay.AlipaySetup()
	if err != nil {
		return
	}
	bizContent := make(map[string]interface{})

	bizContent["out_trade_no"] = userData.OutRequestNo //"outOrderNo123"
	bizContent["out_request_no"] = userData.OutRequestNo
	bizContent["order_title"] = userData.OrderTitle
	bizContent["amount"] = userData.Amount
	bizContent["product_code"] = userData.ProductCode
	bizContent["payee_user_id"] = userData.PayeeUserId
	bizContent["timeout_express"] = userData.TimeoutExpress

	result, err := client.FundAuthOrderAppFreeze(ctx, bizContent)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, result, "")
}

/*Trade
* @Description: 统一收单交易接口
* @param ctx
 */
func Trade(ctx iris.Context) {

	var userData AuthParams
	if err := ctx.ReadJSON(&userData); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	json := SetBizContent(userData)

	client, err := conalipay.AlipaySetup()
	if err != nil {
		return
	}
	result, err := client.TradeCreate(ctx, json)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, result, "")
}

/*SetBizContent
* @Description: 构造初始化入参
 */
func SetBizContent(user AuthParams) map[string]interface{} {
	// 构造传递参数
	bizContent := make(map[string]interface{})

	bizContent["out_trade_no"] = user.OutTradeNo
	bizContent["total_amount"] = user.TotalAmount
	bizContent["subject"] = user.Subject
	bizContent["buyer_id"] = user.BuyerId
	return bizContent
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

/*OrderSync
* @Description: 订单数据同步接口
* @param ctx
* @return {}
 */
func OrderSync(ctx iris.Context) {}
