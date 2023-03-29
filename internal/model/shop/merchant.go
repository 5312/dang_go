package shop

import (
	"dang_go/internal/database"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// 商家
type Merchant struct {
	gorm.Model
	CompanyName             string  `json:"company_name" validate:"required" gorm:"not null;comment:企业名称"`
	LegalRepresentativeName string  `json:"legal_representative_name" validate:"required" gorm:"not null;comment:法定代表人名称"`
	IdNumber                string  `json:"id_number" validate:"required" gorm:"not null;comment:法人身份证号"`
	Phone                   string  `json:"phone" validate:"required" gorm:"not null;comment:联系电话"`
	ConsumerHotline         string  `json:"consumer_hotline" validate:"required" gorm:"not null;comment:客服电话"`
	DetailAddress           string  `json:"detail_address" validate:"required" gorm:"not null;comment:公司详情地址"`
	CreditCode              string  `json:"credit_code" validate:"required" gorm:"not null;comment:统一社会信用代码"`
	UnIdNumberPng           string  `json:"un_id_number_png" validate:"required" gorm:"not null;comment:身份证正反面"`
	BusinessLicense         string  `json:"business_license" validate:"required" gorm:"not null;comment:营业执照"`
	ShopPicture             string  `json:"shop_picture" validate:"required" gorm:"not null;comment:商铺图片"`
	Avatar                  string  `json:"avatar" validate:"required" gorm:"not null;comment:头像"`
	Introduction            string  `json:"introduction" validate:"required" gorm:"not null;comment:简介"`
	City                    string  `json:"city" validate:"required" gorm:"not null;comment:城市"`
	ShopName                string  `json:"shop_name" validate:"required" gorm:"not null;comment:商铺名称"`
	ShopAdminName           string  `json:"shop_admin_name" validate:"required" gorm:"not null;comment:开户的管理员名称"`
	ShopAdminPhone          string  `json:"shop_admin_phone" validate:"required" gorm:"not null;comment:开户的管理员电话"`
	AlipayNumber            string  `json:"alipay_number" validate:"required" gorm:"not null;comment:支付宝账户"`
	AlipayName              string  `json:"alipay_name" validate:"required" gorm:"not null;comment:支付宝账户名称"`
	RentBalance             string  `json:"rent_balance"  gorm:"comment:租金余额"`
	WithdrawalRatio         string  `json:"withdrawal_ratio"  gorm:"comment:提现比例"`
	PurchasePrice           int     `json:"purchase_price"  gorm:"comment:商家采购价"`
	AddressLease            *Medium `json:"address_lease" gorm:"type:json;租赁地址:数组类型的json数据"`
	AddressReturn           *Medium `json:"address_return" gorm:"type:json;归还地址:数组类型的json字符串"`
}

// []存放地址
type Medium struct {
	Promise   string `json:"promise"`
	AvatarUrl string `json:"avatarUrl"`
	City      string `json:"city"`
	Address   string `json:"address"`
	Province  string `json:"province"`
	District  string `json:"district"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
}

// 使用 Validate单例, 缓存结构体信息
var validate *validator.Validate

/*Create 开户 */
func (e *Merchant) AddShop() (id int, err error) {
	validate = validator.New()
	verr := validate.Struct(e)
	if verr != nil {
		fmt.Printf("参数验证失败 %v -- \n", verr)
		err = verr
		return
	}
	result := database.DB.Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetPage 查*/
func (e *Merchant) GetPage() (Merchant []Merchant, err error) {
	table := database.DB.Model(&e)
	//if err = table.Where("name like ?", "").Order("sort").Find(&User).Error; err != nil {
	//	return
	//}
	if err = table.Find(&Merchant).Error; err != nil {
		return
	}
	return
}

// 修改
func (e *Merchant) Update(id uint) (update Merchant, err error) {
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

// ! 重要 为模型实现Value/Scan函数
func (c Medium) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Medium) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 修改租赁地址
func (e *Medium) UpdateAddress(id uint, data Medium) (success bool, err error) {
	var update Merchant

	if err = database.DB.Model(&Merchant{}).Find(&update, id).Error; err != nil {
		fmt.Printf("err--", err)
		success = false
		return
	}
	res, err := json.Marshal(data)
	//fmt.Printf("获取json%v \n ", res)
	if err = database.DB.Model(&Merchant{}).Model(&update).Update("address_lease", res).Error; err != nil {
		fmt.Printf("查询%v \n", err)
		success = false
		return
	}
	success = true
	return
}
