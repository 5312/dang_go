package pay

import (
	"dang_go/internal/model/shop"
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
	ormCr, err := params.Create()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, ormCr, "添加成功")
}
