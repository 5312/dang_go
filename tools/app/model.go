package app

import (
	"fmt"
	"math/rand"
	"time"
)

// Response 列表数据
type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

// Page 分页数据
type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

// PageResponse 分页数据
type PageResponse struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data Page `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

func (r *Response) ReturnOK() *Response {
	r.Code = 200
	return r
}
func (r *PageResponse) ReturnOK() *PageResponse {
	r.Code = 200
	return r
}

func (r *Response) ReturnError(code int) *Response {
	r.Code = code
	return r
}

/*OrderOutBizNo
* @Description: 生成外部订单id
* @return string
 */
func OrderOutBizNo() string {
	// 生成当前时间戳的字符串
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestampStr := fmt.Sprintf("%d", timestamp)

	// 生成 6 位随机数字
	randomNum := rand.Intn(899999) + 100000

	// 组合商家订单号
	orderNum := timestampStr + fmt.Sprintf("%d", randomNum)
	return orderNum
}
