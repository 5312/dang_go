// Package system 会员
package system

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
)

// 会员信息

type Member struct {
	gorm.Model
	Name          string `json:"name" gorm:"not null;comment:会员名称"`
	Reality       string `json:"reality" gorm:"not null;comment:真实姓名"`
	Phone         string `json:"phone" gorm:"not null;comment:手机号"`
	BonusPoints   string `json:"bonus_points" gorm:"not null;comment:奖励分"`
	PromoterId    string `json:"promoter_id" gorm:"not null;comment:推广商id"`
	PromoterManId string `json:"promoter_man_id" gorm:"not null;comment:推广员id"`
	IdNumber      string `json:"id_number" gorm:"not null;comment:身份证号"`
	InflowStatus  string `json:"inflow_status" gorm:"not null;comment:入流状态:1搜索 2平台扫码;"`
}

/*Create 增 */
func (e *Member) Create() (err error) {
	result := database.DB.Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetPage 查*/
func (e *Member) GetPage() (Member []Member, err error) {
	table := database.DB.Model(&e)

	if err = table.Find(&Member).Error; err != nil {
		return
	}
	return
}
