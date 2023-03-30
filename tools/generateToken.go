package tools

import jwt "dang_go/middleware"

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
func GenerateToken(claims jwt.CustomClaims) (LoginResult, error) {
	j := &jwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}

	token, err := j.CreateToken(claims)
	userInfo := map[string]interface{}{
		"id":   claims.ID,
		"name": claims.Name,
	}
	data := LoginResult{
		User:  userInfo,
		Token: token,
	}
	return data, err
}
