package shop

import "gorm.io/gorm"

// Shop
// @Description: 商品管理
type Shop struct {
	gorm.Model
	Picture            string `json:"picture" gorm:"not null;comment:商品图片: json化数组"`
	Name               string `json:"name" gorm:"not null;comment:商品名称"`
	IsSearch           string `json:"is_search" gorm:"not null;comment:是否可被搜索"`
	Title              string `json:"title" gorm:"not null;comment:商品标题"`
	MarketingStatement string `json:"marketing_statement" gorm:"not null;comment:营销语句"`
	Category           string `json:"category" gorm:"comment:商品分类"`
	LongSort           string `json:"long_sort" gorm:"comment:长租短租"`
	ScreenProtection   string `json:"screen_protection" gorm:"not null;comment:碎屏保障"`
	Freight            string `json:"freight" gorm:"comment:运费"`
	OptionalWith       string `json:"optional_with" gorm:"comment:可选搭配"`

	Sort int    `json:"sort" gorm:"comment:排序"`
	Type int    `json:"type" gorm:"comment:类型"`
	Unit string `json:"unit" gorm:"comment:单位"`
}
