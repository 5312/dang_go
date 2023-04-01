package main

import (
	// 本地包
	. "dang_go/api"
	. "dang_go/config"
	. "dang_go/internal/database"
	"dang_go/internal/model/gorm"
	_ "dang_go/pay"
	"fmt"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func main() {
	fmt.Println("|---------------------------|")
	fmt.Println("|----------admin------------|")
	fmt.Println("|---------------------------|")

	// 1. 读取配置config.InitConfig()
	c := DbConfig{}
	c.InitConfig()

	// 2. 初始化数据库
	var db Database
	db = new(Mysql)
	db.InitGormDB()
	// 3. 迁移表
	_ = gorm.AutoMigrate(DB)
	fmt.Printf("数据库结构初始化成功！\n")

	// 4.注册路由
	InitUser()

}
