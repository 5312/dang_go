package controller

import (
	"dang_go/internal/model/system"
	jwt "dang_go/middleware"
	"dang_go/tools"
	"dang_go/tools/app"
	"encoding/json"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-pay/gopay/alipay"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"time"
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

type RequestData struct {
	AlipayUserId string        `json:"alipay_user_id"`
	RefreshToken string        `json:"refresh_token"`
	UserId       string        `json:"user_id"`
	AccessToken  string        `json:"access_token"`
	User         system.Member `json:"user"`
	Token        string        `json:"token"`
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
	//fmt.Printf("%v 序列化 \n", jsonData.Response.Code)
	// TODO: 密钥暂定读取,以后需要加密
	// 换取授权访问令牌（默认使用utf-8，RSA2）
	// appId：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
	//  codeOrToken：支付宝授权码或refresh_token
	var (
		appId       = viper.GetString("Alipay.AppId")   // "2021003186611052" // 租宝贝
		privateKey  = viper.GetString("Alipay.RSA2048") //string(data)
		grantType   = "authorization_code"
		codeOrToken = code.AuthCode
		signType    = "RSA2"
	)

	success, err := alipay.SystemOauthToken(ctx, appId, privateKey, grantType, codeOrToken, signType)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	//fmt.Printf("返回的支付宝信息%v \n", success)
	// 添加数据库
	mem := system.Member{
		Name:          jsonData.Response.NickName,
		Reality:       "",
		Phone:         "",
		BonusPoints:   "",
		PromoterId:    "",
		PromoterManId: "",
		IdNumber:      "",
		InflowStatus:  "",
		ZfbUserId:     success.Response.UserId,
		Avatar:        jsonData.Response.Avatar,
	}
	fmt.Printf("userid: %v \n", success.Response.UserId)
	createErr := mem.Create(success.Response.UserId)

	claims := jwt.CustomClaims{
		ID:       mem.ID,
		Name:     mem.Name,
		Password: mem.ZfbUserId,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 过期时间 一小时
			Issuer:    "admin",                  //签名的发行者
		},
	}
	tokenObj, tokenErr := tools.GenerateToken(claims, mem)
	if tokenErr != nil {
		app.Error(ctx, -1, tokenErr, "")
		return
	}
	requestData := RequestData{
		AlipayUserId: success.Response.AlipayUserId,
		RefreshToken: success.Response.RefreshToken,
		AccessToken:  success.Response.AccessToken,
		UserId:       success.Response.UserId,
		User:         mem,
		Token:        tokenObj.Token,
	}
	if createErr != nil {
		app.Error(ctx, -1, createErr, "")
		return
	}
	app.OK(ctx, requestData, "登录成功")

}
