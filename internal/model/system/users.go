// Package system 账号管理
package system

import (
	"dang_go/internal/database"
	jwt "dang_go/middleware"
	jwtgo "github.com/dgrijalva/jwt-go"
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

type LoginResult struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

/*Login 登录*/
func (e *User) Login(name string, password string) (token LoginResult, err error) {
	var Users []User
	table := database.DB.Model(&e)

	if err = table.Debug().Where("name = ?", name).Where("password = ?", password).Find(&Users).Error; err != nil {
		return
	}

	if len(Users) <= 0 {
		// 没有用户
		err = errors.New("用户名不存在")
		return
	}
	generateTokens, err := generateToken(Users[0])
	return generateTokens, nil
}

// generateToken 生成令牌  创建jwt风格的token
func generateToken(user User) (LoginResult, error) {
	j := &jwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}
	claims := jwt.CustomClaims{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效时间
			ExpiresAt: time.Now().Unix() + 86400, // 过期时间6 *  6 * 24 24小时
			Issuer:    "admin",                   //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	userInfo := map[string]interface{}{
		"id":      user.ID,
		"name":    user.Name,
		"age":     user.Age,
		"email":   user.Email,
		"account": user.Account,
	}
	data := LoginResult{
		User:  userInfo,
		Token: token,
	}
	return data, err
}
