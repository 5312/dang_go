package database

import (
	"gorm.io/gorm"
)

type Database interface {
	Open(dialector gorm.Dialector, conn gorm.Option) (db *gorm.DB, err error)
	GetConnect() string
}
