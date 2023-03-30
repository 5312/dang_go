package shop

import (
	"dang_go/internal/database"
	jwt "dang_go/middleware"
	"dang_go/tools"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gogf/gf/util/gconv"
	"github.com/kataras/iris/v12/x/errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

// Merchant 商家
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
	PassWord                string  `json:"pass_word"  gorm:"not null;comment:密码默认电话后六位" `
	AlipayNumber            string  `json:"alipay_number" validate:"required" gorm:"not null;comment:支付宝账户"`
	AlipayName              string  `json:"alipay_name" validate:"required" gorm:"not null;comment:支付宝账户名称"`
	RentBalance             string  `json:"rent_balance"  gorm:"comment:租金余额"`
	WithdrawalRatio         string  `json:"withdrawal_ratio"  gorm:"comment:提现比例"`
	PurchasePrice           int     `json:"purchase_price"  gorm:"comment:商家采购价"`
	AddressLease            *Medium `json:"address_lease" gorm:"type:json;租赁地址:数组类型的json数据"`
	AddressReturn           *Medium `json:"address_return" gorm:"type:json;归还地址:数组类型的json字符串"`
}
type MerchantInfo struct {
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
	RentBalance             string `json:"rent_balance"  gorm:"comment:租金余额"`
	WithdrawalRatio         string `json:"withdrawal_ratio"  gorm:"comment:提现比例"`
	PurchasePrice           int    `json:"purchase_price"  gorm:"comment:商家采购价"`
}

// Medium []存放地址
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

//goland:noinspection GoUnusedParameter
func (e *Merchant) Login(name string, password string) (token tools.LoginResult, err error) {
	var shangjia []Merchant
	table := database.DB.Model(&e)

	if err = table.Debug().Where("shop_admin_phone = ?", name).Where("pass_word = ?", password).Find(&shangjia).Error; err != nil {

		return
	}

	if len(shangjia) <= 0 {
		// 没有用户
		err = errors.New("用户名不存在")
		return
	}

	// 构造 CustomClaims 对象
	claims := jwt.CustomClaims{
		ID:       shangjia[0].ID,
		Name:     shangjia[0].ShopAdminPhone,
		Password: shangjia[0].PassWord,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效时间
			ExpiresAt: time.Now().Unix() + 86400, // 过期时间6 *  6 * 24 24小时
			Issuer:    "admin",                   //签名的发行者
		},
	}
	var userInfo MerchantInfo
	errgconv := gconv.Struct(shangjia[0], &userInfo)
	if errgconv != nil {
		fmt.Printf("结构体转化失败%v", errgconv)
	}
	generateTokens, err := tools.GenerateToken(claims, userInfo)

	return generateTokens, nil
}

/*AddShop
* @Description: 开户
* @receiver e Merchant
* @return id 返回id
* @return err 错误处理
 */
func (e *Merchant) AddShop() (id int, err error) {
	fmt.Printf("参数 %v \n", e.PassWord)
	if e.PassWord == "" {
		e.PassWord = e.getLastSixDigits()
	}

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

/*GetLastSixDigits
* @Description: 获取后6位密码
* @receiver u
* @return string
 */
func (e *Merchant) getLastSixDigits() string {
	if len(e.ShopAdminPhone) >= 6 {
		return strings.TrimSpace(e.ShopAdminPhone[len(e.ShopAdminPhone)-6:])
	}
	return ""
}

/*GetPage
* @Description:查询商户列表
* @receiver *Merchant
* @return Merchant
* @return err
 */
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

/*Update
* @Description: 修改商户信息
* @receiver e
* @param id
* @return update
* @return err
 */
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

/*Value
* @Description: ! 重要 为模型实现Value/Scan函数
* @receiver c 方法接收者
* @return driver.Value
* @return error
 */
func (e *Medium) Value() (driver.Value, error) {
	b, err := json.Marshal(e)
	return string(b), err
}
func (e *Medium) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), e)
}

/*UpdateAddress
* @Description: 添加修改租赁地址
* @receiver e
* @param id
* @param data
* @param updateType
* @return success
* @return err
 */
func (e *Medium) UpdateAddress(id uint, data Medium, updateType string) (success bool, err error) {
	var update Merchant
	result := database.DB.Model(&Merchant{}).Find(&update, id)
	if err = result.Error; err != nil {
		//fmt.Printf("err--", err)
		success = false
		return
	}
	if result.RowsAffected == 0 {
		success = false
		err = errors.New("找不到对应记录")
		return
	}
	// 转为 切片类型
	//mediums := []Medium{
	//	data,
	//}
	res, err := json.Marshal(data)
	fmt.Printf("获取json%v \n ", result.RowsAffected == 0)
	if err = database.DB.Model(&Merchant{}).Model(&update).Update(updateType, res).Error; err != nil {
		fmt.Printf("查询%v \n", err)
		success = false
		return
	}
	success = true
	return
}
