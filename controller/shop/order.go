package shop

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

//"平台端：商品订单"

// 查询自己分配的订单列表

/*PayOrder
* @Description: 订单支付
* @param ctx
 */
func PayOrder(ctx iris.Context) {
	oid, _ := ctx.Params().GetUint("OID")
	rid, _ := ctx.Params().GetUint("RID")

	fmt.Printf("oid%v \n", oid)
	fmt.Printf("rid:%v \n", rid)

}
