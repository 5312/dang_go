package app

import (
	"github.com/kataras/iris/v12"
)

// 声明map 错误类型
var ErrorText = map[string]string{
	"record not found": "找不到记录",
	"EOF":              "请输入参数",
	"response is nil or access_token is NULL": "access_token is Null", // 确认appid 是否对应
	"WHERE conditions required":               "条件不存在",
}

/*OK 成功数据处理*/
func OK(c iris.Context, data interface{}, msg string) {

	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	err := c.JSON(res.ReturnOK())
	if err != nil {
		return
	}
}

/*Error 失败数据处理*/
func Error(c iris.Context, code int, err error, msg string) {
	var res Response
	//res.Msg = err.Error()
	if res.Msg = ErrorText[err.Error()]; res.Msg == "" {
		res.Msg = err.Error()
	}

	if msg != "" {
		res.Msg = msg
	}
	errors := c.JSON(res.ReturnError(code))
	if errors != nil {
		return
	}
}
