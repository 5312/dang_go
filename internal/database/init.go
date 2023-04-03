package database

import (
	"bytes"
	"dang_go/tools/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var DB *gorm.DB

func (e *Mysql) Setup() *sql.DB {
	var db Database
	db = new(Mysql)
	dsn := db.GetConnect()

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	orm, err := db.Open(mysql.New(mysqlConfig), &gorm.Config{})

	DB = orm

	if err != nil {
		fmt.Print("\n", "db--", db)
		fmt.Print("\n", "-----------------")
		fmt.Print("\n", "err", err)
		fmt.Print("\n", "链接数据库失败")
		panic(err)
	}
	// 链接池
	sqldb, err := orm.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqldb.SetConnMaxLifetime(time.Hour)

	if err = sqldb.Ping(); err != nil {
		panic(err)
	}

	return sqldb
}

type Mysql struct{}

func (e *Mysql) GetConnect() string {
	// 读取db配置
	c := config.GetConfig()

	var conn bytes.Buffer
	conn.WriteString(c.User)
	conn.WriteString(":")
	conn.WriteString(c.Password)
	conn.WriteString("@tcp(")
	conn.WriteString(c.Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(c.Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(c.DriverName)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=10000ms")
	return conn.String()
}

func (e *Mysql) Open(dialector gorm.Dialector, conn gorm.Option) (db *gorm.DB, err error) {
	return gorm.Open(dialector, conn)
}
