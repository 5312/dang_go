package system

import (
	"dang_go/internal/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12/x/errors"
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
func (e *User) GetPage(name string) (User []User, err error) {
	table := database.DB.Model(&e)
	var n = "%" + name + "%"
	if err = table.Where("name like ?", n).Order("sort").Find(&User).Error; err != nil {
		return
	}
	return
}

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

/*Login 登录*/
func (e *User) Login(name string, password string) (User []User, err error) {
	table := database.DB.Model(&e)

	if err = table.Debug().Where("name = ?", name).Where("password = ?", password).Find(&User).Error; err != nil {
		return
	}

	if len(User) <= 0 {
		// 没有用户
		err = errors.New("用户名不存在")
		return
	}

	return
}
