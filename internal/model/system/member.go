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
	Reality       string `json:"reality" gorm:"comment:真实姓名"`
	Phone         string `json:"phone" gorm:"comment:手机号"`
	BonusPoints   string `json:"bonus_points" gorm:"comment:奖励分"`
	PromoterId    string `json:"promoter_id" gorm:"comment:推广商id"`
	PromoterManId string `json:"promoter_man_id" gorm:"comment:推广员id"`
	IdNumber      string `json:"id_number" gorm:"comment:身份证号"`
	IdNumberUrl   string `json:"id_number_url" gorm:"comment:身份证号照片"`
	IsStudent     string `json:"is_student" gorm:"comment:是否学生"`
	OrderCount    string `json:"order_count" gorm:"comment:订单信息"`
	InflowStatus  string `json:"inflow_status" gorm:"comment:入流状态:1搜索 2平台扫码;"`
	ZfbUserId     string `json:"zfb_user_id" gorm:"not null;comment:支付宝user_id"`
	Avatar        string `json:"avatar" gorm:"comment:头像"`
	Lng           string `json:"lng" gorm:"comment:下单时经度"`
	Lat           string `json:"lat" gorm:"comment:下单时纬度"`
}

/*Create 增 */
func (e *Member) Create(ID string) (id uint, err error) {
	// 先查询
	var Id []Member
	database.DB.Model(&e).Where("zfb_user_id = ?", ID).Find(&Id)
	if len(Id) != 0 {
		//fmt.Printf("会员已添加不重复添加", Id[0].ID)
		id = Id[0].ID
		return
	}
	result := database.DB.Create(&e)
	//fmt.Printf("擦汗如%v", e.ID)
	id = e.ID
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

/*Update 改 */
func (e *Member) Update(id uint) (update Member, err error) {
	table := database.DB.Model(&e)

	result := table.Where("id = ?", id).Updates(&e)

	if err = result.Error; err != nil {
		return
	}

	if err = table.First(&update, id).Error; err != nil {
		return
	}

	return
}
