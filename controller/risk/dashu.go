package risk

import (
	"dang_go/tools/aes"
	"dang_go/tools/app"
	"encoding/base64"
	"encoding/json"
	"github.com/kataras/iris/v12"
)

type UtilsFun struct {
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	Ident_number  string  `json:"ident_number"`
	Service       string  `json:"service"`
	Mode          string  `json:"mode"`
	Goods_type    int     `json:"goods_type"`
	Status        int     `json:"status"`
	Total_rent    float64 `json:"total_rent"`
	Total_periods int     `json:"total_periods"`
	Price         float64 `json:"price"`
}

/*TestDashu
* @Description: 天狼星测试
* @param ctx
 */
func TestDashu(ctx iris.Context) {

	// 接收参数
	var data UtilsFun
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "请输入参数")
		return
	}
	key := "Q0ymUIe1t26ZfG7s"

	encryptData := aes.Encrypt(data, key)
	encryptDataWith64 := base64.StdEncoding.EncodeToString(encryptData)
	//fmt.Println("encryptData: ", encryptData)
	//fmt.Println("encryptDataWith64: ", encryptDataWith64)
	// 测试环境url
	url := "http://sit-shouwei.shouxin168.com/sandbox/lightning/product/query"
	institutionId := "d6518f1e-9270-11e9-890c-9801a79f5a77"

	body, err := TianLangFengKong(encryptDataWith64, url, institutionId)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	var v interface{}
	json.Unmarshal(body, &v)
	res := v.(map[string]interface{})
	app.OK(ctx, res, "查询成功")

}

/*Dashu
* @Description: 天狼星生产
* @param ctx
 */
func Dashu(ctx iris.Context) {

	// 接收参数
	var data UtilsFun
	if err := ctx.ReadJSON(&data); err != nil {
		app.Error(ctx, -1, err, "请输入参数")
		return
	}
	url := "https://shouwei.shouxin168.com/api/lightning/product/query"
	institutionId := "90495182887569066062508527805056"
	key := "BgT6aQDZsck2cdnr"

	encryptData := aes.Encrypt(data, key)
	encryptDataWith64 := base64.StdEncoding.EncodeToString(encryptData)

	body, err := TianLangFengKong(encryptDataWith64, url, institutionId)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	var v interface{}
	json.Unmarshal(body, &v)
	res := v.(map[string]interface{})
	app.OK(ctx, res, "查询成功")

}
