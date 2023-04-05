package shop

import (
	"dang_go/internal/database"
	"dang_go/tools/app"
	"gorm.io/gorm"
)

// Order
// @Description:订单
type Order struct {
	gorm.Model
	FormMerchant    uint   `json:"form_merchant" gorm:"comment:所属商家"`
	FormMember      uint   `json:"form_member" gorm:"comment:所属会员"`
	Name            string `json:"name" gorm:"not null;comment:商品名称"`
	Picture         string `json:"picture" gorm:"comment:商品图片"`
	MaterialId      string `json:"material_id" gorm:"not null;comment:素材 ID"`
	MaterialKey     string `json:"material_key" gorm:"not null;comment:素材标识"`
	OrderOutBizNo   string `json:"order_out_biz_no" gorm:"comment:外部id"`
	Address         string `json:"address" gorm:"comment:地址"`
	City            string `json:"city" gorm:"comment:城市"`
	Province        string `json:"province" gorm:"comment:行省"`
	District        string `json:"district" gorm:"comment:区域"`
	Params          string `json:"params" gorm:"type:varchar(2048);comment:参数"`
	Phone           string `json:"phone" gorm:"comment:电话"`
	CurrentAddress  string `json:"currentAddress" gorm:"comment:目前的地址"`
	Status          string `json:"status" gorm:"comment:订单状态"`
	Message         string `json:"message" gorm:"comment:用户留言"`
	Mark            string `json:"mark" gorm:"comment:备注"`
	RiskInformation string `json:"risk_information" gorm:"comment:风控理由"`
	RiskGrade       string `json:"risk_grade" gorm:"comment:风控等级"`
}

/*Create
* @Description: 数据库增加记录
 */
func (o *Order) Create(memberId uint) (id int, err error) {
	o.OrderOutBizNo = app.OrderOutBizNo()
	o.FormMember = memberId
	result := database.DB.Create(&o)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetPage
* @Description: 会员获取自己订单
* @receiver o
* @param id
* @return list
* @return err
 */
func (o *Order) GetPage(id uint) (list []Order, err error) {
	table := database.DB.Model(&o)

	res := table.Where("form_member = ?", id).Find(&list)
	if res.Error != nil {
		err = res.Error
		return
	}
	return
}

/*GetCount
* @Description: 商家获取累计订单数
* @receiver o
* @param id
* @return num
* @return err
 */
func (o *Order) GetCount(id uint) (num int64, err error) {
	table := database.DB.Model(&o)
	res := table.Where("form_merchant = ?", id).Count(&num)
	if res.Error != nil {
		err = res.Error
		return
	}
	return
}

/*GetMerchantOrder
* @Description: 商家获取订单列表
* @receiver o
* @param id
* @return list
* @return err
 */
func (o *Order) GetMerchantOrder(id uint) (list []Order, err error) {
	table := database.DB.Model(&o)
	res := table.Where("form_merchant = ?", id).Find(&list)
	if res.Error != nil {
		err = res.Error
		return
	}
	return
}

/*GetOrderDetail
* @Description: 订单详情
* @receiver o
* @return {}
 */
func (o *Order) GetOrderDetail(id uint) (list Order, err error) {
	table := database.DB.Model(&o)

	res := table.Debug().Where("id = ?", id).Find(&list)
	if res.Error != nil {
		err = res.Error
		return
	}
	return
}

/*GetAllPage
* @Description: 全部订单
* @receiver o
* @param status
* @return list
* @return err
 */
func (o *Order) GetAllPage(status string) (list []Order, err error) {
	table := database.DB.Model(&o)
	if status == "" {
		res := table.Find(&list)
		if res.Error != nil {
			return
		}
		return
	}
	res := table.Where("status = ?", status).Find(&list)
	if res.Error != nil {
		err = res.Error
		return
	}
	return

}
