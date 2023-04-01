package applets

import (
	"dang_go/internal/database"
	"dang_go/internal/model/shop"
	"github.com/kataras/iris/v12/x/errors"
)

// RecommendProduct
// @Description: 推荐 从商品继承
type RecommendProduct struct {
	shop.Shop
}

/*AddFromShopList
* @Description: 从商品 添加
* @receiver s
* @return id
* @return err
 */
func (r *RecommendProduct) AddFromShopList(formId string) (id uint, err error) {
	var shops *shop.Shop
	table := database.DB.Model(&shops)

	var list RecommendProduct
	getData := table.Where("id = ?", formId).Find(&list)
	if getData.RowsAffected == 0 {
		err = errors.New("商品不存在")
		return
	}
	result := database.DB.Create(&list)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetRecommend
* @Description: 获取活动商品
* @receiver r
 */
func (r *RecommendProduct) GetRecommend() (list []RecommendProduct, err error) {
	table := database.DB.Model(&r)
	if err = table.Find(&list).Error; err != nil {
		return
	}
	return
}
