package main

import (
	// 本地包
	. "dang_go/api"
	. "dang_go/config"
	. "dang_go/internal/database"
	"dang_go/internal/model/gorm"
	"fmt"
)

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
	// 4.注册路由
	InitUser()

}
