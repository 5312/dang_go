package tools

import (
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/client"
)

var restClient *client.RestClient

func init() {
	//var err error
	// configFilePath 是 SDK 初始化的配置文件，格式为 JSON 串，具体配置信息见后面
	//configFilePath := "/tmp/rest-config.json"
	//restClient, err = client.NewRestClient(configFilePath) // 初始化客户端
	//if err != nil {
	//	panic(fmt.Errorf("failed to NewRestClient err:%+v", err))
	//}
}

/** 以下是 configFilePath 的示例文件
BizId：链的 ID
RestUrl：rest 的服务地址
AccessId：分配给用户用于访问 rest 的账户名（可通过开放联盟链控制台获得）
AccessSecret：分配给用于用户访问rest的密钥的路径（可通过开放联盟链控制台获得）
MaxIdleConns：最大空闲连接数
IdleConnTimeout：空闲连接超时时间，单位为秒
RetryMaxAttempts：请求重试最大次数
BackOffPeriod：重试时间间隔，单位为毫秒
**/
/**
{
 "RestUrl": "https://rest.baas.alipay.com",
 "AccessId": "baas_admin",
 "AccessSecret": "/Users/Documents/go_workspace/src/gitlab.alipay-inc.com/antchain/restclient-go-sdk/test/access.key",
 "MaxIdleConns": 10,
 "IdleConnTimeout": 30,
 "RetryMaxAttempts": 5,
 "BackOffPeriod": 500
}
**/
