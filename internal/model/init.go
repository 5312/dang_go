package model

import (
	"database/sql"
	"fmt"
	"time"

	"com.example.dang/config"
	"com.example.dang/internal/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB() *sql.DB {
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
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})

	DB = db

	if err != nil {
		fmt.Print("\n", "db", db)
		fmt.Print("\n", "-----------------")
		fmt.Print("\n", "err", err)
		fmt.Print("\n", "链接数据库失败")
		panic(err)
	}
	// 链接池
	sqldb, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqldb.SetConnMaxLifetime(time.Hour)

	if err = sqldb.Ping(); err != nil {
		panic(err)
	}

	//自动迁移表
	tableInit(db)

	return sqldb
}

func tableInit(db *gorm.DB) {
	// 菜单表
	// 创建表时添加后缀
	db.AutoMigrate(&system.Menu{})

}
