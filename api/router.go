package api

import (
	"dang_go/api/v1/sys"

	"github.com/kataras/iris/v12"
)

// InitUser 初始化api接口
func InitUser(app *iris.Application) {
	// V1
	v1 := app.Party("/v1")
	sys.RegisterRoute(v1)
}
