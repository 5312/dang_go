package shop

import "gorm.io/gorm"

// Bill
// @Description: 账单
type Bill struct {
	gorm.Model
	Address           string        `json:"address" gorm:"comment:地址"`
	City              string        `json:"city" gorm:"comment:城市"`
	ClearDate         string        `json:"clearDate" gorm:"comment:"`
	ClosingOrderTime  string        `json:"closingOrderTime" gorm:"comment:"`
	DeliverTime       string        `json:"deliverTime" gorm:"comment:"`
	District          string        `json:"district" gorm:"comment:区县"`
	FlowState         int           `json:"flowState" gorm:"comment:状态"`
	Name              string        `json:"name" gorm:"comment:用户信息"`
	OrderCode         string        `json:"orderCode" gorm:"comment:订单ID"`
	OutTradeNo        string        `json:"outTradeNo" gorm:"comment:"`
	Payment           int           `json:"payment" gorm:"comment:"`
	PaymentRecordList []interface{} `json:"paymentRecordList" gorm:"comment:"`
	Period            int           `json:"period" gorm:"comment:"`
	PeriodCount       string        `json:"periodCount" gorm:"comment:" `
	Price             int           `json:"price" gorm:"comment:"`
	ProductParams     string        `json:"productParams" gorm:"comment:"`
	Province          string        `json:"province" gorm:"comment:"`
	Ratio             int           `json:"ratio" gorm:"comment:"`
	RefundDate        string        `json:"refundDate" gorm:"comment:"`
	RentPrice         string        `json:"rentPrice" gorm:"comment:"`
	RName             string        `json:"r_name" gorm:"comment:"`
	RPhone            string        `json:"r_phone" gorm:"comment:"`
	Tit               string        `json:"tit" gorm:"comment:"`
	TradeNo           string        `json:"tradeNo" gorm:"comment:"`
}
