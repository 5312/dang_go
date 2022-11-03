package model

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatasource() (err error) {
	// 读取db配置
	user := viper.Get("DB.User")
	password := viper.Get("DB.Password")
	host := viper.Get("DB.Host")
	port := viper.Get("DB.Port")
	driverName := viper.Get("DB.DriverName")
	charset := viper.Get("DB.Charset")

	dsn := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s", user, ":", password, "@tcp(", host, ":", port, ")", "/", driverName, "?charset=", charset, "&parseTime=True&loc=Local")
	// fmt.Printf("dsn: %v\n", dsn)
	// 连接云服务器 mysql
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
	// 关闭链接
	//defer sqlDB.Close()
	return
}

func tableInit(db *gorm.DB) {
	// ims_sys_log 日志表
	// db.AutoMigrate(&models.TakeSysLog{})
}
