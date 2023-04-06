package tools

import (
	jwt "dang_go/middleware"
)

type LoginResult struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

/*GenerateToken
* @Description: 生成令牌  创建jwt风格的token
* @param claims
* @return LoginResult
* @return error
 */
func GenerateToken(claims jwt.CustomClaims, users interface{}) (LoginResult, error) {
	j := &jwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}

	token, err := j.CreateToken(claims)

	data := LoginResult{
		User:  users,
		Token: token,
	}
	return data, err
}
