// Package promoter
// @Description: 推广员
package promoter

import (
	"dang_go/internal/database"
	"dang_go/middleware"
	"github.com/kataras/iris/v12/x/errors"
	"gorm.io/gorm"
)

// Personnel
// @Description: 推广员
type Personnel struct {
	gorm.Model
	FromPromoter    uint   `json:"from_promoter" gorm:"comment:所属推广商id"`
	Name            string `json:"name" gorm:"comment:姓名"`
	Phone           string `json:"phone" gorm:"comment:手机号"`
	Settlement      string `json:"settlement" gorm:"comment:结算类型|按比例|按订单"`
	Amount          string `json:"amount" gorm:"comment:结算金额/比例"`
	PersonnelClient int    `json:"personnel_client" gorm:"comment:客户数"`
	MyOrder         int    `json:"my_order" gorm:"comment:推广订单"`
	ToPromoteLink   int    `json:"to_promote_link" gorm:"comment:推广链接"`
	AccountBalance  int    `json:"account_balance" gorm:"comment:账户余额"`
}

/*Create 增 */
func (p *Personnel) Create() (id uint, err error) {
	result := database.DB.Create(&p)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*Delete 删除 */
func (p *Personnel) Delete(id uint) (err error) {
	table := database.DB.Model(&p)
	//fmt.Printf("%v", p.FromPromoter)
	result := table.Where("from_promoter = ?", p.FromPromoter).Where("id = ?", id).Delete(&Personnel{})
	if result.RowsAffected == 0 {
		err = errors.New("数据不存在")
		return
	}
	return
}

/*Update 改 */
func (p *Personnel) Update(id uint) (update Personnel, err error) {
	if err = database.DB.Model(&p).First(&update, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = database.DB.Model(&p).Model(&update).Updates(&p).Error; err != nil {
		return
	}
	return
}

/*GetMyPersonnelList
* @Description: 查询
* @receiver p
* @param userInfo
* @return list
* @return err
 */
func (p *Personnel) GetMyPersonnelList(userInfo *middleware.CustomClaims) (list []Personnel, err error) {
	table := database.DB.Model(&p)
	//fmt.Printf("userinof -- %v", userInfo.ID)
	if err = table.Where("from_promoter = ?", userInfo.ID).Find(&list).Error; err != nil {
		return
	}

	return
}
