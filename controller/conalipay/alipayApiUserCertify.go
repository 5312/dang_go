// Package conalipay
// @Description: 身份验证
package conalipay

import (
	"dang_go/tools/app"
	"fmt"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"strings"
	"time"
)

/*AlipaySetup
* @Description: 初始化 支付宝 client
* @return client
 */
func AlipaySetup() (client *alipay.Client, err error) {
	// 初始化支付宝客户端
	//	appid：应用ID
	//	privateKey：应用私钥，支持PKCS1和PKCS8
	//	isProd：是否是正式环境
	appId := viper.GetString("Alipay.AppId")
	client, err = alipay.NewClient(appId, viper.GetString("Alipay.RSA2048"), viper.GetBool("Alipay.isProd"))

	if err != nil {
		xlog.Error(err)
	}
	// 打开Debug开关，输出日志，默认关闭
	//client.DebugSwitch = gopay.DebugOn

	client.SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
							SetNotifyUrl("https://www.fmm.ink") // 设置异步通知URL
	return
}

/*DecryptPhoneNum
* @Description: 手机号解密
* @param ctx
 */
func DecryptPhoneNum(ctx iris.Context) {
	type Decry struct {
		EncryptedData string
	}
	var phones Decry
	if err := ctx.ReadJSON(&phones); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	// 解密支付宝开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    secretKey:AES密钥，支付宝管理平台配置
	//    beanPtr:需要解析到的结构体指针
	var secretKey = "GPiEZ1pCJXPeAy1iiMjaAw=="

	phone := new(alipay.UserPhone)
	err := alipay.DecryptOpenDataToStruct(phones.EncryptedData, secretKey, phone)
	fmt.Printf("phoone%v", err)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	app.OK(ctx, phone, "")
}

type AuthParams struct {
	IdCard string `json:"idCard"`
	Name   string `json:"name"`
}

/*AlipayUserCertifyOpenInitializeRequest
* @Description:身份认证初始化服务接口
* @return {}
 */
func AlipayUserCertifyOpenInitializeRequest(ctx iris.Context) {
	var userData AuthParams
	if err := ctx.ReadJSON(&userData); err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	json := SetBizContent(userData)
	// 初始化 支付宝
	client, _ := AlipaySetup()

	// 身份认证初始化服务接口
	result, err := client.UserCertifyOpenInit(ctx, json) //client.UserCertifyOpenCertify(ctx, json)
	//fmt.Printf("身份验证成功 %v \n", result)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	// 开始身份认证
	certify := make(map[string]interface{})
	certify["certify_id"] = result.Response.CertifyId

	certifyUrl, initErr := client.UserCertifyOpenCertify(ctx, certify)

	urlAndId := make(map[string]interface{})

	urlAndId["certifyId"] = result.Response.CertifyId
	urlAndId["url"] = certifyUrl
	if initErr != nil {
		app.Error(ctx, -1, initErr, "")
		return
	}
	app.OK(ctx, urlAndId, "")
}

/*SetBizContent
* @Description: 构造初始化入参
 */
func SetBizContent(user AuthParams) map[string]interface{} {

	identityObj := make(map[string]interface{})
	identityObj["identity_type"] = "CERT_INFO"
	identityObj["cert_type"] = "IDENTITY_CARD"
	identityObj["cert_name"] = user.Name
	identityObj["cert_no"] = user.IdCard
	//构造商户配置json对象
	merchantConfigObj := make(map[string]interface{})
	//// 设置回调地址
	merchantConfigObj["return_url"] = "/"

	// 构造传递参数
	bizContent := make(map[string]interface{})
	// 商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合;建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	outerOrderNo := fmt.Sprintf("yrzj%d%s", time.Now().Unix(), uuid.New().String())
	outerOrderNo = strings.Replace(outerOrderNo, "-", "", -1)[:31]

	bizContent["outer_order_no"] = outerOrderNo
	bizContent["biz_code"] = "FACE"
	bizContent["identity_param"] = identityObj
	bizContent["merchant_config"] = merchantConfigObj

	return bizContent
}
