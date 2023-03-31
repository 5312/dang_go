// Package system 账号管理
package system

import (
	"dang_go/internal/database"
	jwt "dang_go/middleware"
	"dang_go/tools"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/util/gconv"
	"github.com/kataras/iris/v12/x/errors"
	"gorm.io/gorm"
	"time"
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
type UserInfo struct {
	Name    string  `json:"name" gorm:"not null;comment:姓名"`
	Age     uint8   `json:"age" gorm:"comment:年龄"`
	Email   *string `json:"email" gorm:"comment:邮箱"`
	Account int     `json:"account" gorm:"not null;comment:账号"`
	Sort    int     `json:"sort" gorm:"comment:显示顺序;"`
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

/*Login
* @Description: 用户名登录 TODO: 账号登录
* @receiver e
* @param name
* @param password
* @return token
* @return err
 */
func (e *User) Login(name string, password string) (token tools.LoginResult, err error) {
	var Users []User
	table := database.DB.Model(&e)

	if err = table.Where("name = ?", name).Where("password = ?", password).Find(&Users).Error; err != nil {
		return
	}

	if len(Users) <= 0 {
		// 没有用户
		err = errors.New("用户不存在")
		return
	}
	// 构造 CustomClaims 对象
	claims := jwt.CustomClaims{
		ID:       Users[0].ID,
		Name:     Users[0].Name,
		Password: Users[0].Password,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效时间
			ExpiresAt: time.Now().Unix() + 86400, // 过期时间6 *  6 * 24 24小时
			Issuer:    "admin",                   //签名的发行者
		},
	}
	var userInfo UserInfo
	errgconv := gconv.Struct(Users[0], &userInfo)
	if err != nil {
		fmt.Printf("转换失败%v", errgconv)
	}

	generateTokens, err := tools.GenerateToken(claims, userInfo)
	return generateTokens, nil
}
