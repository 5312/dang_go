package database

import "database/sql"

type Database interface {
	InitGormDB() (db *sql.DB)
}
