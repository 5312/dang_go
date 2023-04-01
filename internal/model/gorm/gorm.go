package gorm

import (
	"dang_go/internal/model/applets"
	"dang_go/internal/model/promoter"
	"dang_go/internal/model/shop"
	"dang_go/internal/model/system"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		new(system.Menu),
		new(system.User),
		new(system.Member), // 会员
		new(shop.Merchant), // 商家
		new(shop.Shop),     // 上坪
		new(shop.Category), // 分类

		new(promoter.Promoter),  // 推广商
		new(promoter.Personnel), // 推广员
		// 小程序
		new(applets.HomeImg), // banner图
	)
}
