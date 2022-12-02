package system

import (
	"dang_go/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name" gorm:"not null;comment:姓名"`
	Age      uint8   `json:"age" gorm:"comment:年龄"`
	Email    *string `json:"email" gorm:"comment:邮箱"`
	Account  int     `json:"account" gorm:"not null;comment:账号"`
	Password string  `json:"password" gorm:"not null;comment:密码"`
	Sort     int     `json:"sort" gorm:"comment:显示顺序;"`
}

/*Create 增 */
func (e *User) Create() (id int, err error) {
	result := database.DB.Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

/*GetPage 查*/
func (e *User) GetPage(n string) (User []User, err error) {
	table := database.DB.Model(&e)

	if err = table.Where("name like ?", "%@name% ", map[string]interface{}{
		"name": n,
	}).Order("sort").Find(&User).Error; err != nil {
		return
	}
	return
}
