package middleware

import (
	"dang_go/tools/app"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// JWTAuth 中间件，检查token
func JWTAuth(ctx iris.Context) {

	token := ctx.Request().Header.Get("Authorization")

	if token == "" {
		app.Error(ctx, -1, errors.New("Token is expired"), "请求未携带token，无权限访问")
		ctx.StopExecution()
		return
	}
	//fmt.Print("get token: ", token)
	j := NewJWT()
	// parseToken 解析token包含的信息
	// 将获取的Authorization 内容通过分割出来
	authorArr := strings.SplitN(token, " ", 2)
	// Authorization的字符串通常是 "Bearer" 开头(可以理解为固定格式,标识使用承载模式),然后一个空格 再加上token的内容
	// Tips:  请求头中Authorization的内容直接是token也是可以的
	if len(authorArr) != 2 || authorArr[0] != "Bearer" {
		app.Error(ctx, -1, errors.New("Bearer"), "request header Authorization formal error")
		return
	}
	claims, err := j.ParseToken(authorArr[1])

	if err != nil {
		if err == TokenExpired {
			app.Error(ctx, -1, errors.New("授权已过期"), "授权已过期")
			ctx.StopExecution()
			return
		}
		app.Error(ctx, -1, err, "")
		ctx.StopExecution()
		return
	}
	// 继续交由下一个路由处理,并将解析出的信息传递下去
	ctx.Values().Set("claims", claims)
	ctx.Next()
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID       uint   `json:"userId"`
	Name     string `json:"name"`
	Password string `json:"telephone"`
	jwt.StandardClaims
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

//// 这是SignKey
//func SetSignKey(key string) string {
//	SignKey = key
//	return SignKey
//}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		return j.SigningKey, nil
	})
	//fmt.Printf("签名: %v \n", token)
	//fmt.Printf("签名err: %v \n", err)

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
