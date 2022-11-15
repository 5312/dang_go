package system

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null" `
	Path       string `json:"path" gorm:"not null;comment:api路径" `
	Component  string `json:"component" gorm:"not null;comment:组件路径" `
	Parent_id  uint   `json:"parent_id" gorm:"not null;comment:父id" `
	Icon       string `json:"icon" gorm:"comment:icon图标" `
	Permission string `json:"permission" gorm:"comment:权限" `
	Note       string `json:"note" gorm:"comment:备注"  `
	Type       int    `json:"type" gorm:"comment:类型:0菜单1节点2权限" `
	Status     int8   `json:"status" gorm:"comment:状态:1正常 2禁用" `
	Sort       int    `json:"sort" gorm:"comment:显示顺序" `
	Hide       bool   `json:"hide" gorm:"comment:显示:0否1" `
}
