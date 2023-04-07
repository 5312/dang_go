package system

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
)

// MemAddress
// @Description: 收获地址
type MemAddress struct {
	gorm.Model
	FormMember uint   `json:"form_member" gorm:"所属会员"`
	Name       string `json:"name" gorm:"comment:收货人"`
	Phone      string `json:"phone" gorm:"comment:手机号"`
	///province city district
	Province string `json:"province" gorm:"comment:省"`
	City     string `json:"city" gorm:"comment:市"`
	District string `json:"district" gorm:"comment:区"`
	Address  string `json:"address" gorm:"comment:详细地址"`
	Def      bool   `json:"def" gorm:"comment:默认地址"`
}

/*Create 增 */
func (m *MemAddress) Create(ID uint) (id uint, err error) {
	m.FormMember = ID
	result := database.DB.Create(&m)
	id = m.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetList
* @Description: 获取收获地址
* @receiver m
* @return list
* @return err
 */
func (m *MemAddress) GetList(memberId uint) (list []MemAddress, err error) {
	table := database.DB.Model(&m)
	res := table.Where("form_member = ?", memberId).Find(&list)
	if res.Error != nil {
		err = res.Error
		return
	}
	return
}
