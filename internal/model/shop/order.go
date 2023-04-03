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
	Name           string `json:"name" gorm:"not null;comment:商品名称"`
	Picture        string `json:"picture" gorm:"comment:商品图片"`
	MaterialId     string `json:"material_id" gorm:"not null;comment:素材 ID"`
	MaterialKey    string `json:"material_key" gorm:"not null;comment:素材标识"`
	OrderOutBizNo  string `json:"order_out_biz_no" gorm:"comment:外部id"`
	Address        string `json:"address" gorm:"comment:地址"`
	City           string `json:"city" gorm:"comment:城市"`
	Province       string `json:"province" gorm:"comment:行省"`
	District       string `json:"district" gorm:"comment:区域"`
	Params         string `json:"params" gorm:"type:varchar(2048);comment:参数"`
	Phone          string `json:"phone" gorm:"comment:电话"`
	CurrentAddress string `json:"currentAddress" gorm:"comment:目前的地址"`
}

/*Create
* @Description: 数据库增加记录
 */
func (o *Order) Create() (id int, err error) {
	o.OrderOutBizNo = app.OrderOutBizNo()
	result := database.DB.Create(&o)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
