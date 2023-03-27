package shop

import (
	"dang_go/internal/database"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	CompanyName             string `json:"company_name" validate:"required" gorm:"not null;comment:企业名称"`
	LegalRepresentativeName string `json:"legal_representative_name" validate:"required" gorm:"not null;comment:法定代表人名称"`
	IdNumber                string `json:"id_number" validate:"required" gorm:"not null;comment:法人身份证号"`
	Phone                   string `json:"phone" validate:"required" gorm:"not null;comment:联系电话"`
	ConsumerHotline         string `json:"consumer_hotline" validate:"required" gorm:"not null;comment:客服电话"`
	DetailAddress           string `json:"detail_address" validate:"required" gorm:"not null;comment:公司详情地址"`
	CreditCode              string `json:"credit_code" validate:"required" gorm:"not null;comment:统一社会信用代码"`
	UnIdNumberPng           string `json:"un_id_number_png" validate:"required" gorm:"not null;comment:身份证正反面"`
	BusinessLicense         string `json:"business_license" validate:"required" gorm:"not null;comment:营业执照"`
	ShopPicture             string `json:"shop_picture" validate:"required" gorm:"not null;comment:商铺图片"`
	Avatar                  string `json:"avatar" validate:"required" gorm:"not null;comment:头像"`
	Introduction            string `json:"introduction" validate:"required" gorm:"not null;comment:简介"`
	City                    string `json:"city" validate:"required" gorm:"not null;comment:城市"`
	ShopName                string `json:"shop_name" validate:"required" gorm:"not null;comment:商铺名称"`
	ShopAdminName           string `json:"shop_admin_name" validate:"required" gorm:"not null;comment:开户的管理员名称"`
	ShopAdminPhone          string `json:"shop_admin_phone" validate:"required" gorm:"not null;comment:开户的管理员电话"`
	AlipayNumber            string `json:"alipay_number" validate:"required" gorm:"not null;comment:支付宝账户"`
	AlipayName              string `json:"alipay_name" validate:"required" gorm:"not null;comment:支付宝账户名称"`
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
