package system

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null;"`
	Path       string `json:"path" gorm:"not null;comment:api路径;"`
	Component  string `json:"component" gorm:"not null;comment:组件路径;" `
	ParentId   uint   `json:"parentId" gorm:"type:int(11);column:parent_id;comment:父id"`
	Icon       string `json:"icon" gorm:"comment:icon图标;"`
	Permission string `json:"permission" gorm:"comment:权限;" `
	Note       string `json:"note" gorm:"comment:备注;"`
	Type       int    `json:"type" gorm:"comment:类型:0菜单1节点2权限;"`
	Status     int8   `json:"status" gorm:"comment:状态:1正常 2禁用;" `
	Sort       int    `json:"sort" gorm:"comment:显示顺序;"`
	NoCache    bool   `json:"noCache" gorm:"column:no_cache"`
}
type Menus struct {
	gorm.Model
	Name       string  `json:"name" gorm:"not null;"`
	Path       string  `json:"path" gorm:"not null;comment:api路径;"`
	Component  string  `json:"component" gorm:"not null;comment:组件路径;" `
	ParentId   uint    `json:"parentId" gorm:"type:int(11);column:parent_id;comment:父id"`
	Icon       string  `json:"icon" gorm:"comment:icon图标;"`
	Permission string  `json:"permission" gorm:"comment:权限;" `
	Note       string  `json:"note" gorm:"comment:备注;"`
	Type       int     `json:"type" gorm:"comment:类型:0菜单1节点2权限;"`
	Status     int8    `json:"status" gorm:"comment:状态:1正常 2禁用;" `
	Sort       int     `json:"sort" gorm:"comment:显示顺序;"`
	NoCache    bool    `json:"noCache" gorm:"column:no_cache"`
	Children   []Menus `json:"children"`
}

/*Create 增 */
func (e *Menu) Create() (id int, err error) {
	result := database.DB.Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*Delete 删除 */
func (e *Menu) Delete(id uint) (success bool, err error) {
	// .Where("id = ?", id)
	if err = database.DB.Model(&e).Where("id = ?", id).Delete(&Menu{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

/*Update 改 */
func (e *Menu) Update(id uint) (update Menu, err error) {
	if err = database.DB.Model(&e).First(&update, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = database.DB.Model(&e).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

/*GetPage 查*/
func (e *Menu) GetPage() (Menus []Menu, err error) {
	table := database.DB.Model(&e)

	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}
