package main

import (
	"fmt"

	"com.example.dang/api"
	"com.example.dang/config"
	"com.example.dang/internal/model"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("|---------------------------|")
	fmt.Println("|----------admin------------|")
	fmt.Println("|---------------------------|")
	/* 初始化iris */
	app := iris.Default()

	// config.InitConfig()
	c := config.DbConfig{}
	c.InitConfig()

	// 初始化数据库
	model.InitGormDB()

	// 注册路由 中间件
	api.InitUser(app)

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
