package model

import (
	"fmt"
	"time"

	"com.example.dang/config"
	"com.example.dang/internal/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormDB() *gorm.DB {
	// 读取db配置
	c := config.GetConfig()

	dsn := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s", c.User, ":", c.Password, "@tcp(", c.Host, ":", c.Port, ")", "/", c.DriverName, "?charset=", c.Charset, "&parseTime=True&loc=Local")

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	// fmt.Printf("dsn: %v\n", dsn)
	// 连接云服务器 mysql
	DB, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})

	if err != nil {
		fmt.Print("\n", "db", DB)
		fmt.Print("\n", "-----------------")
		fmt.Print("\n", "err", err)
		fmt.Print("\n", "链接数据库失败")
		panic(err)
	}
	// 链接池
	sqlDB, err := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}

	//自动迁移表
	tableInit(DB)

	return DB
}

func tableInit(db *gorm.DB) {
	// 菜单表
	db.AutoMigrate(&system.Menus{})
}
