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
//
//goland:noinspection GoErrorStringFormat,GoErrorStringFormat,GoErrorStringFormat,GoErrorStringFormat
var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
	SignKey          = "newtrekWang"
)

// JWT
// @Description: JWT 签名结构
type JWT struct {
	SigningKey []byte
}

/*NewJWT
* @Description: 新建一个jwt实例
* @return *JWT
 */
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

/*JWTAuth
* @Description: JWTAuth 中间件，检查token
* @param ctx
 */
func JWTAuth(ctx iris.Context) {

	if LoginNoAuth(ctx) {
		ctx.Next()
		return
	}

	token := ctx.Request().Header.Get("Authorization")

	if token == "" {
		app.Error(ctx, -1, errors.New("token is expired"), "请求未携带token，无权限访问")
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
		//goland:noinspection GoErrorStringFormat
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

// CustomClaims
// @Description: 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID                 uint   `json:"userId"`
	Name               string `json:"name"`
	Password           string `json:"telephone"`
	jwt.StandardClaims        // { 过期时间 签发者 生效时间}
}

/*GetSignKey
* @Description: 获取signKey
* @return string
 */
func GetSignKey() string {
	return SignKey
}

/*CreateToken
* @Description:  CreateToken 生成一个token
* @receiver j
* @param claims
* @return string
* @return error
 */
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

/*ParseToken
* @Description:  解析Tokne
* @receiver j
* @param tokenString
* @return *CustomClaims
* @return error
 */
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

/*RefreshToken
* @Description: 更新token
* @receiver j
* @param tokenString
* @return string
* @return error
 */
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
