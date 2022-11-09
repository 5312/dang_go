package main

import (
	"fmt"
	// 本地包
	. "dang_go/api"
	. "dang_go/config"
	. "dang_go/internal/model"
	"dang_go/middleware"

	// 第三方
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("|---------------------------|")
	fmt.Println("|----------admin------------|")
	fmt.Println("|---------------------------|")
	/* 初始化iris */
	app := iris.Default()

	/* 中间件 */
	// app.Use(middleware.Cors) // 跨域 !!! 失效
	app.UseGlobal(middleware.Cors)

	// config.InitConfig()
	c := DbConfig{}
	c.InitConfig()

	// 初始化数据库
	InitGormDB()

	// 注册路由 中间件
	InitUser(app)

	// 5 注册view  默认静态页面
	tmpl := iris.HTML("./web/static", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	// 读取conf
	str := ":"
	port := viper.Get("Conf.Port")
	res := fmt.Sprintf("%s%s", str, port)
	app.Run(iris.Addr(res))

}
