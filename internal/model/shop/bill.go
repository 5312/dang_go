package shop

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
)

// Bills
// @Description:账单
type Bills struct {
	gorm.Model
	Address           string `json:"address" gorm:"comment:地址"`
	City              string `json:"city" gorm:"comment:城市"`
	ClosingOrderTime  string `json:"closing_order_time" gorm:"comment:订单关闭时间"`
	District          string `json:"district" gorm:"comment:区县"`
	FlowState         int    `json:"flow_state" gorm:"comment:状态"`
	Name              string `json:"name" gorm:"comment:用户信息"`
	OrderCode         string `json:"order_code" gorm:"comment:订单ID"`
	OutTradeNo        string `json:"out_trade_no" gorm:"comment:外部订单号"`
	Payment           int    `json:"payment" gorm:"comment:支付"`
	PaymentRecordList string `json:"payment_record_list" gorm:"comment:支付"`
	Period            int    `json:"period" gorm:"comment:"`
	PeriodCount       string `json:"period_count" gorm:"comment:" `
	Price             int    `json:"price" gorm:"comment:金额"`
	ProductParams     string `json:"product_params" gorm:"comment:参数"`
	Province          string `json:"province" gorm:"comment:-"`
	Ratio             int    `json:"ratio" gorm:"comment:-"`
	RentPrice         string `json:"rent_price" gorm:"comment:-"`
	RName             string `json:"r_name" gorm:"comment:-"`
	RPhone            string `json:"r_phone" gorm:"comment:-"`
	Tit               string `json:"tit" gorm:"comment:-"`
	TradeNo           string `json:"trade_no" gorm:"comment:-"`
}

/*Create
* @Description: 生成账单
 */
func (b *Bills) Create(memberId uint) (id int, err error) {
	//o.OrderOutBizNo = app.OrderOutBizNo()
	//o.FormMember = memberId
	result := database.DB.Create(&b)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetAll
* @Description: 获取全部账单
* @receiver b
* @return list
* @return err
 */
func (b *Bills) GetAll() (list []Bills, err error) {

	result := database.DB.Model(&b).Find(&list)
	if result.Error != nil {
		err = result.Error
		return
	}
	return

}
