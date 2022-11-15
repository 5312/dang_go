package system

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name       string `json:"name" gorm:"comment:菜单name"`
	Parent_id  uint   `json:"parent_id" gorm:"comment:父id"`
	Icon       string `json:"icon" gorm:"comment:icon图标"`
	Path       string `json:"path" gorm:"comment:api路径"`
	Component  string `json:"component" gorm:"comment:组件径"`
	Permission string `json:"permission" gorm:"comment:权限"`
	Note       string `json:"note" gorm:"comment:备注"`
	Hide       b ool  `json:"hide" gorm:"comment:显示:0否1"`
	HideInMenu bool  `json:"hideInMenu" gorm:"在菜单中隐藏自己和子节点"`
	Type       *int8 `json:"type" gorm:"comment:类型:0菜单节点2权限"`
	tatus      *int8 `json:"status" gorm:"comment:状态:1正常 2禁用"`
	Sort       *int  `json:"sort" gorm:"comment:显示顺序"`
	Show       bool  `json:"show"`
}
