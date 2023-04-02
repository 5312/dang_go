package shop

import (
	"dang_go/internal/database"
	"dang_go/middleware"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// 使用 Validate单例, 缓存结构体信息
var validate *validator.Validate

// Shop
// @Description: 商品管理
type Shop struct {
	gorm.Model
	FromShops          uint   `json:"form_shops" validate:"required" gorm:"所属商家"`
	Title              string `json:"title" validate:"required" gorm:"not null;comment:商品标题"`
	MarketingStatement string `json:"marketing_statement"  validate:"required" gorm:"not null;comment:营销语句"`
	Category           int    `json:"category" gorm:"comment:商品分类:全新二手"`
	LongSort           int    `json:"long_sort" gorm:"comment:长租短租"`
	ScreenProtection   string `json:"screen_protection"  validate:"required"  gorm:"not null;comment:碎屏保障;"`
	Freight            string `json:"freight" gorm:"comment:运费;"`
	OptionalWith       string `json:"optional_with" gorm:"comment:可选搭配;"`
	OptionalWithImg    string `json:"optional_with_img" gorm:"comment:可选搭配图片;"`
	// 属性类目
	ObjectCategory string `json:"object_category"  validate:"required"  gorm:"not null;comment:类目选择;"`
	Pattern        string `json:"pattern"  validate:"required"  gorm:"not null;comment:型号;"`
	Network        string `json:"network"  validate:"required"  gorm:"not null;comment:网络制式;"`
	Screen         string `json:"screen"   validate:"required"  gorm:"not null;comment:屏幕尺寸;"`
	Version        string `json:"version"  validate:"required"  gorm:"not null;comment:版本;"`
	// 销售信息
	Combo             string `json:"combo"  validate:"required"  gorm:"not null;comment:套餐;"`
	LeasePeriod       string `json:"lease_period" gorm:"not null;comment:租期;"`
	Color             string `json:"color"  validate:"required"  gorm:"not null;comment:颜色[];"`
	MemoryInformation string `json:"memory_information"  validate:"required"  gorm:"not null;comment:运行内存+机身内存;"`
	//商品信息
	//ShopImg    string `json:"shop_img" gorm:"comment:商品图片路径;"`
	Picture    string `json:"picture" validate:"required" gorm:"not null;comment:商品图片: json化数组"`
	ShopDetail string `json:"shop_detail" gorm:"comment:商品详情;"`
	// 支付信息
	IsPreSale bool `json:"is_pre_sale" gorm:"comment:是否预售;"`
	IsSearch  bool `json:"is_search" gorm:"not null;comment:是否可被搜索"`
	//物流信息
	LeaseAddress  string `json:"lease_address"  validate:"required"  gorm:"comment:发货地址;"`
	ReturnAddress string `json:"return_address"  validate:"required"  gorm:"comment:归还地址;"`
	// 其他
	Sort   int    `json:"sort" gorm:"comment:排序"`
	Type   int    `json:"type" gorm:"comment:类型"`
	Unit   string `json:"unit" gorm:"comment:单位"`
	Status string `json:"status" gorm:"comment:商品状态:已上架|审核中|已下架|未通过|草稿箱"`
}

/*AddLeaseShop
* @Description: 添加租赁商品
* @receiver s
 */
func (s *Shop) AddLeaseShop() (id int, err error) {

	validate = validator.New()
	verr := validate.Struct(s)
	if verr != nil {
		fmt.Printf("参数验证失败 %v -- \n", verr)
		err = verr
		return
	}
	result := database.DB.Create(&s)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetMyShopList
* @Description: 获取自己的商品列表
* @receiver s
* @param userInfo
* @return list
* @return err
 */
func (s *Shop) GetMyShopList(userInfo *middleware.CustomClaims) (list []Shop, err error) {
	table := database.DB.Model(&s)
	fmt.Printf("userinof -- %v", userInfo.ID)
	if err = table.Where("from_shops = ?", userInfo.ID).Order("sort").Find(&list).Error; err != nil {
		return
	}

	return
}

/*GetCategoryShopList
* @Description: 根据分类id 获取商品列表
* @receiver s
 */
func (s *Shop) GetCategoryShopList(cateId string) (list []Shop, err error) {
	table := database.DB.Model(&s)
	if err = table.Debug().Where("object_category = ?", cateId).Order("sort").Find(&list).Error; err != nil {
		return
	}
	return
}

/*GetDetail
* @Description:获取商品详情
* @receiver s
* @param shopId
* @return list
* @return err
 */
func (s *Shop) GetDetail(shopId uint) (list Shop, err error) {
	table := database.DB.Model(&s)
	if err = table.Debug().Where("id = ?", shopId).Find(&list).Error; err != nil {
		return
	}
	return
}
