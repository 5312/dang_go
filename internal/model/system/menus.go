package system

import (
	"gorm.io/gorm"
)

type Menus struct {
	gorm.Model
	Name       string
	Code       string
	Parent_id  int    `json:"parent_id" gorm:"comment:父id"`
	Icon       string `json:"icon" gorm:"comment:icon图标"`
	Path       string `json:"path" gorm:"comment:api路径"`
	Component  string `json:"component" gorm:"comment:组件路径"`
	Permission string `json:"permission" gorm:"comment:权限"`
	Type       string `json:"type" gorm:"comment:类型：0菜单 1节点"`
	Status     string `json:"status" gorm:"comment:状态：1正常 2禁用"`
	Hide       string `json:"hide" gorm:"comment:是否可见：1是 2否"`
	Note       string `json:"note" gorm:"comment:备注"`
	Sort       int    `json:"sort" gorm:"comment:显示顺序"`
}
