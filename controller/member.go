package controller

import (
	"dang_go/internal/model/system"
	"dang_go/middleware"
	"dang_go/tools/app"
	"github.com/kataras/iris/v12"
)

/*GetPage
* @Description: 获取会员信息
* @param ctx
 */
func GetPage(ctx iris.Context) {
	// Get all records
	var data system.Member

	result, err := data.GetPage()
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	app.OK(ctx, result, "查询成功")
}

type IdParams struct {
	IdCard string
	Name   string
}

/*SaveIdCard
* @Description: 保存身份证
* @param ctx
 */
func SaveIdCard(ctx iris.Context) {
	var p IdParams
	if err := ctx.ReadJSON(&p); err != nil {
		//panic(err.Error())
		app.Error(ctx, -1, err, "")
		return
	}
	var data system.Member
	//登录用户
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	data.Reality = p.Name
	data.IdNumber = p.IdCard
	//保存为登录用户数据
	result, err := data.Update(userInfo.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "保存成功")
}

type LngLat struct {
	Latitude  string
	Longitude string
}

/*SaveUpLngLat
* @Description: 保存下单时经纬度
* @param ctx
 */
func SaveUpLngLat(ctx iris.Context) {
	var params LngLat
	if err := ctx.ReadJSON(&params); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}

	var data system.Member
	//登录用户
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	data.Lng = params.Longitude
	data.Lat = params.Latitude

	result, err := data.Update(userInfo.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, result, "保存成功")
}

/*Address
* @Description: 添加地址
* @param ctx
 */
func Address(ctx iris.Context) {
	// 获取参数
	var params system.MemAddress
	if err := ctx.ReadJSON(&params); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	// 插入数据
	res, err := params.Create(userInfo.ID)

	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, res, "保存成功")
}

/*AddressList
* @Description: 获取地址
* @param ctx
 */
func AddressList(ctx iris.Context) {
	var address system.MemAddress
	userInfo := ctx.Values().Get("claims").(*middleware.CustomClaims)
	res, err := address.GetList(userInfo.ID)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, res, "")
}
