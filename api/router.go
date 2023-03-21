package api

import (
	"dang_go/api/v1/sys"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	_ "dang_go/docs"
)

// InitSysRouter 初始化api接口
func InitSysRouter(app *iris.Application) {
	// V1
	v1 := app.Party("/v1")
	{
		// sys
		sys.RegisterMenuRoute(v1)
	}
	/* swagger文档*/
	// 指向swagger init生成文档的路径
	//config := &swagger.Config{
	//	URL:         "http://localhost:87/docs/swagger.json",
	//	DeepLinking: true,
	//}
	//app.Get("/swagger/*any", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
	app.Get("/swagger/index.html", swagger.WrapHandler(swaggerFiles.Handler))

	//app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
}
