package shop

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
	"strings"
)

// MerchantMan
// @Description: 成员
type MerchantMan struct {
	gorm.Model
	Name       string `json:"name" gorm:"comment:用户名"`
	Phone      string `json:"phone" gorm:"comment:手机号"`
	MerchantId uint   `json:"merchant_id" gorm:"comment:所属商家"`
	PassWord   string `json:"pass_word" gorm:"comment:密码"`
}

/*Create
* @Description: 添加成员
* @receiver m
 */
func (m *MerchantMan) Create(merchantId uint) (id uint, err error) {
	table := database.DB.Model(&m)
	m.MerchantId = merchantId
	m.PassWord = m.getLastSixDigits()
	result := table.Create(&m)
	id = m.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetLastSixDigits
* @Description: 获取后6位密码
* @receiver u
* @return string
 */
func (m *MerchantMan) getLastSixDigits() string {
	if len(m.Phone) >= 6 {
		return strings.TrimSpace(m.Phone[len(m.Phone)-6:])
	}
	return ""
}

/*GetList
* @Description: 获取商家用户
* @receiver m
* @param merchantId
* @return list
* @return err
 */
func (m *MerchantMan) GetList(merchantId uint) (list []MerchantMan, err error) {
	table := database.DB.Model(&m)
	result := table.Where("merchant_id = ?", merchantId).Find(&list)

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
