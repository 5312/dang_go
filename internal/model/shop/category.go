package shop

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;"`
	ParentId uint   `json:"parentId" gorm:"type:int(11);column:parent_id;comment:父id"`
	Image    string `json:"image" gorm:"comment:海报;"`
	Sort     int    `json:"sort" gorm:"comment:显示顺序;"`
	Type     int    `json:"type" gorm:"comment:类型;"`
}

type Categorys struct {
	gorm.Model
	Name     string      `json:"name" gorm:"not null;"`
	ParentId uint        `json:"parentId" gorm:"type:int(11);column:parent_id;comment:父id"`
	Image    string      `json:"image" gorm:"comment:海报;"`
	Sort     int         `json:"sort" gorm:"comment:显示顺序;"`
	Type     int         `json:"type" gorm:"comment:类型;"`
	Children []Categorys `json:"children"`
}

/*Create 增 */
func (e *Category) Create() (id int, err error) {
	result := database.DB.Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 查寻
func (e *Category) GetPage(n string, s string, end string) (Category []Category, err error) {
	table := database.DB.Model(&e)

	if err = table.Where(`name like @name OR created_at > @startTime AND created_at < @endTime`, map[string]interface{}{
		"name":      "%" + n + "%",
		"startTime": s,
		"endTime":   end,
	}).Order("sort").Find(&Category).Error; err != nil {
		return
	}
	return
}

// 修改
func (e *Category) Update(id uint) (update Category, err error) {
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

// 删除
func (e *Category) Delete(id uint) (success bool, err error) {
	if err = database.DB.Model(&e).Where("id = ?", id).Delete(&Category{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
