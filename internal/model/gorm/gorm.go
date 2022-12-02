package gorm

import (
	"dang_go/internal/model/system"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		new(system.Menu),
		new(system.User),
	)
}
