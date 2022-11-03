package api

import (
	"com.example.dang/api/v1/sys"
	"github.com/kataras/iris/v12"
)

// 初始化api接口
func InitUser(app *iris.Application) {
	sys.RegisterRoute(app)
}
