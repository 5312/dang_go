package controller

import (
	"dang_go/internal/model/system"
	"dang_go/tools/app"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay/alipay"
	"github.com/kataras/iris/v12"
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
	var (
		appId       = "2021001188609380"
		privateKey  = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCLdEHC+cU2WkuM+pmordmUVpSAEitw+vLUOo+Suh8s+iKESsKFOX6CRf7Vl16agelsf5rHzDd2RPT0NicTT0Dg7n6L+qxco/M5D8JWYAhLDdy/57bExesBNGyrXiDM+E5cCxJQFBePX2Im9t6csQNwOQ20vtFe0uIKIUL5YJpoGmEkNvf8wNUkMB5ewWJlL1CvqZWWLIu4CFIIM7IC+zkxXLygLsq7/Wu/Pi9EIg+LVZp5OrZeMKAsafkHnxvfANgVS3B4wUX9PDgMKnMpDZGrXpOJ16HSrOhxqvhXJjLoq+vAmkqfKMgC0w2+IvLNitPNIuNPv/+QOvn2sJQELh7VAgMBAAECggEAFLTPaubxTG+91hIDrNMbgnYUmKuZ/a2VTsPfO5cdN+1wIJqFJYjy6t7/xrEyH9j4Ut7jmZUOCyaUBIrh5HEZHgzrh1dSLnM9kxRu86pQsOw/AdOjBLaqfUROemeAkU6xO+N89Zz2Hpi4FzcCuNtjuk8OZO8MYXsIix6DNkoMwzYnV5xsgire/PJV5AuBQ+y22tijXTqo3eMXbZNWRhitlPqPJqzSMmQZ8SbqKnJc7JPZTN8iJUK6nAc8MB2P7/e9KiS33RXN1Mh52M2kuoCH8DsklU4ZKYe+pIwZRel+5pez+NJ+WLy+PCC7P10kaDvaBlnQVGbdjuvTjHIKmTgT4QKBgQDGOAG5XVm+W5nSrf1QmgA+449BMdR/K2eTWWJLYNlAlnO0f8EYpaG/Y1F+2mYzMXtd8NaSepUeZF7bjsASlS4VRspPZDp35JiHDSXfV9nFH+01dCVvwLm56yqWT1i7hM1vHYcagKk+VpFdwmoODN1qH0ywUZ162VyCrd7aD+G3HQKBgQC0GvZ/ZyWa585IXZ+L2ZnbsaaSAg7OUTntGS+jk6osppZRJo8VXe72EQfHFKrNubk/9csbpPjgVmy1/q27+4OLiY4ZRaAdgPD3fk7NvizQ3YwFpyIfWUx941JgItOgleqtxBQu5yT8zHVEVGvloyCHRRVzhQ/QkKJC+kZ6i8KhGQKBgFrtMCVG5DcFL//L3mrN6hTvMDS41gBr+bxHAWcQizgsi/EdtYdH23W+6pBlQQJ0zWGfa96Pqr3hTv4qcoNTuWr7KSzYDlYXH7y71EaqtvPNHHQrzkyAPPJDJSsERPDoD5DMG8CVio1VCqPW3e2KCzt2Fii/l5zV/rDXb2XaaTElAoGBALImwZEAPIhfpwKZ6UlycuiHb1aZxn96hSvsb90EIZ6NIb8fvwTJp6eq7OCVpuZcQcvsm326z5tIobvcMYnyngoWhIKnBlxowPJu9BA7fyUTIIAu9GfB8xHLHB6QFHmSEVLU04oZhdKxg4WVZC0AERr80N01z5DkoDxUSAbL3DdxAoGAZza4ac5hTWDTLUVrtSNkJWTKDbt/4XDF2iWcnd/OXXA1gswIYY7aVu7iMacouFqEOYsFWcGT7upB05e/vltLdU3FVCuGb2h37QwBXHRXQpMShyDdIJpVqIGNbSnqUAuLex1p67B27QIQ3nADzq1ztTAhmfGnFM0JFgUJUxfAPHY="
		grantType   = "authorization_code"
		codeOrToken = code.AuthCode
		signType    = "RSA2"
	)
	success, err := alipay.SystemOauthToken(ctx, appId, privateKey, grantType, codeOrToken, signType)
	if err != nil {
		app.Error(ctx, -1, err, "登录失败")
		return
	}
	type RequestData struct {
		UserId      string
		AccessToken string
	}
	//fmt.Printf("%v \n", success)
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
		}

	createErr := mem.Create()
	if createErr != nil {
		app.Error(ctx, -1, createErr, "")
		return
	}
	app.OK(ctx, requestData, "登录成功")
	//app.OK(ctx, success, "登录成功")

}
