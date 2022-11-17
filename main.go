package main

import (
	// 本地包
	. "dang_go/api"
	. "dang_go/config"
	. "dang_go/internal/database"
	"dang_go/internal/model/gorm"
	. "dang_go/middleware"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("|---------------------------|")
	fmt.Println("|----------admin------------|")
	fmt.Println("|---------------------------|")
	/* 1. 初始化iris */
	app := iris.New()

	/* 2. 中间件 */
	// app.Use(middleware.Cors) // 跨域 !!! 失效
	app.Use(Cors)

	// 3. 读取配置config.InitConfig()
	c := DbConfig{}
	c.InitConfig()

	// 4. 初始化数据库
	var db Database
	db = new(Mysql)
	db.InitGormDB()
	// 5. 迁移表
	_ = gorm.AutoMigrate(DB)

	// 6.注册路由 中间件
	InitUser(app)

	// 7 注册view  默认静态页面
	tmpl := iris.HTML("./web/static", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	// 读取conf
	str := ":"
	port := viper.Get("Conf.Port")
	res := fmt.Sprintf("%s%s", str, port)
	_ = app.Run(iris.Addr(res))

}
