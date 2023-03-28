// Package promoter 推广商管理
package promoter

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
)

type Promoter struct {
	gorm.Model
	Name             string `json:"name" gorm:"not null;comment:推广商名称"`
	AdminName        string `json:"admin_name"   gorm:"not null;comment:管理员姓名"`
	AdminPhoneNumber string `json:"admin_phone_number"   gorm:"not null;comment:管理员手机号"`
	AdminPassword    string `json:"admin_password"   gorm:"not null;comment:管理员密码"`
	Settlement       string `json:"settlement"   gorm:"not null;comment:结算类型: 1按比例结算2按订单结算"`
	AccountType      string `json:"account_type"   gorm:"not null;comment:提现账户类型"`
	Alipay           string `json:"alipay"   gorm:"not null;comment:支付宝"`
	AlipayName       string `json:"alipay_name"   gorm:"not null;comment:支付宝姓名"`
}

/*Create 增 */
func (e *Promoter) Create() (id uint, err error) {
	result := database.DB.Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetPage 查*/
func (e *Promoter) GetPage(name string) (Promoter []Promoter, err error) {
	table := database.DB.Model(&e)
	var n = "%" + name + "%"
	if err = table.Where("name like ?", n).Find(&Promoter).Error; err != nil {
		return
	}
	return
}

/*Delete 删除 */
func (e *Promoter) Delete(id uint) (err error) {
	table := database.DB.Model(&e)
	if err = table.Where("id = ?", id).Delete(&Promoter{}).Error; err != nil {
		return
	}
	return
}

/*Update 改 */
func (e *Promoter) Update(id uint) (update Promoter, err error) {
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
