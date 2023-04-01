package conalipay

import (
	"dang_go/internal/model/applets"
	"dang_go/tools/app"
	"fmt"
	"github.com/kataras/iris/v12"
)

/*AddBanner
* @Description:添加banner
* @param ctx
 */
func AddBanner(ctx iris.Context) {
	var data applets.HomeImg
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

/*GetBanner
* @Description:获取数据
* @param ctx
 */
func GetBanner(ctx iris.Context) {
	// Get all records
	var data applets.HomeImg
	types := ctx.URLParam("type") //Params().GetInt("type")
	result, err := data.GetBannerList(types)
	fmt.Printf("types %v", types)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, result, "查询成功")
}

/*UpdateBanner
* @Description:修改banner
* @param ctx
 */
func UpdateBanner(ctx iris.Context) {
	// 接收参数
	var update applets.HomeImg
	id, _ := ctx.Params().GetUint("ID")
	if err := ctx.ReadJSON(&update); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	// 接受参数
	result, err := update.Update(id)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "修改成功")
}

/*AddRecommend
* @Description: 添加商品至活动
* @param ctx
* @return {}
 */
func AddRecommend(ctx iris.Context) {
	var addData applets.RecommendProduct
	id := ctx.URLParam("shopid")

	result, err := addData.AddFromShopList(id)

	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "修改成功")
}

/*GetRecommend
* @Description: 获取活动商品
* @param ctx
* @return {}
 */
func GetRecommend(ctx iris.Context) {
	var addData applets.RecommendProduct
	result, err := addData.GetRecommend()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "修改成功")
}
