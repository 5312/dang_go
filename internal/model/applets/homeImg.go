package applets

import (
	"dang_go/internal/database"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// 使用 Validate单例, 缓存结构体信息
var validate *validator.Validate

// HomeImg
// @Description: // 首页 的展示 图
type HomeImg struct {
	gorm.Model
	ImagePath    string `json:"image_path" validate:"required" gorm:"comment:图片地址"`
	Show         bool   `json:"show"  gorm:"comment:显隐"`
	Name         string `json:"name" validate:"required" gorm:"comment:名称"`
	Price        int    `json:"price" validate:"required" gorm:"comment:价格"`
	Link         string `json:"link" validate:"required" gorm:"comment:链接"`
	OutAppId     string `json:"out_app_id" gorm:"comment:外部appid"`
	OutViewRoute string `json:"out_view_route" gorm:"comment:外部页面路径"`
	Type         int    `json:"type"  validate:"required,oneof=1 2 3 4 5 6" gorm:"not null;comment:图片所属类型:1banner 2category 3排列图 4 5 6 为iphone 13 11 12 行"`
}

/*Create 增 */
func (h *HomeImg) Create() (id uint, err error) {
	validate = validator.New()
	verr := validate.Struct(h)
	if verr != nil {
		fmt.Printf("参数验证失败 %v -- \n", verr)
		err = verr
		return
	}

	result := database.DB.Create(&h)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*Update 改 */
func (h *HomeImg) Update(id uint) (update HomeImg, err error) {
	if err = database.DB.Model(&h).First(&update, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = database.DB.Model(&h).Model(&update).Updates(&h).Error; err != nil {
		return
	}
	return
}

/*GetBannerList
* @Description: 查询
* @receiver p
* @param userInfo
* @return list
* @return err
 */
func (h *HomeImg) GetBannerList(types string) (list []HomeImg, err error) {
	table := database.DB.Model(&h)

	if types == "" {
		if err = table.Find(&list).Error; err != nil {
			return
		}
		return
	}
	if err = table.Debug().Where("type = ?", types).Find(&list).Error; err != nil {
		return
	}
	return
}
