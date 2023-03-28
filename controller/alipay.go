package controller

import (
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay/alipay"
	"github.com/kataras/iris/v12"
	"os"
)

type Response struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	Avatar      string `json:"avatar"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Gender      string `json:"gender"`
	NickName    string `json:"nickName"`
}

type Result struct {
	AuthCode string `json:"authCode"`
	Param    string `json:"param"`
}

type ResponseWrapper struct {
	Response Response `json:"response"`
}

type ParamWrapper struct {
	ResponseWrapper ResponseWrapper `json:"response"`
}

// AlipayLogin 支付宝登录,保存用户信息  , 处理返回信息
func AlipayLogin(ctx iris.Context) {
	// 接收参数
	var code Result
	if err := ctx.ReadJSON(&code); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	response := code.Param
	var jsonData ResponseWrapper
	err := json.Unmarshal([]byte(response), &jsonData)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	fmt.Printf("%v 序列化 \n", jsonData.Response.Code)
	// 换取授权访问令牌（默认使用utf-8，RSA2）
	// appId：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
	//  codeOrToken：支付宝授权码或refresh_token
	// 读取私钥文件
	data, err := os.ReadFile("./config/应用私钥RSA2048-敏感数据，请妥善保管.txt")
	if err != nil {
		fmt.Println("读取私钥文件时发生错误：", err)
		return
	}
	var (
		appId       = "2021003183685933"
		privateKey  = string(data)
		grantType   = "authorization_code"
		codeOrToken = code.AuthCode
		signType    = "RSA2"
	)
	//alipay.UserInfoShare()
	// TODO: user_id 返回失败
	success, err := alipay.SystemOauthToken(ctx, appId, privateKey, grantType, codeOrToken, signType)
	if err != nil {

		app.Error(ctx, -1, err, "")
		return
	}
	type RequestData struct {
		UserId      string
		AccessToken string
	}
	fmt.Printf("返回的支付宝信息%v \n", success)
	requestData := RequestData{
		UserId:      success.Response.UserId,
		AccessToken: success.Response.AccessToken,
	}
	// 添加数据库
	mem :=
		system.Member{
			Name:          jsonData.Response.NickName,
			Reality:       "",
			Phone:         "",
			BonusPoints:   "",
			PromoterId:    "",
			PromoterManId: "",
			IdNumber:      "",
			InflowStatus:  "",
			ZfbUserId:     success.Response.UserId,
		}
	createErr := mem.Create(success.Response.UserId)
	if createErr != nil {
		app.Error(ctx, -1, createErr, "")
		return
	}
	app.OK(ctx, requestData, "登录成功")
	//app.OK(ctx, success, "登录成功")

}
