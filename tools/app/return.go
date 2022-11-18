package app

import (
	"github.com/kataras/iris/v12"
)

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
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	errors := c.JSON(res.ReturnError(code))
	if errors != nil {
		return
	}
}
