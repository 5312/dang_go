package middleware

import (
	"github.com/kataras/iris/v12"
	"strings"
)

func InitMiddleware(r *iris.Application) {
	// 跨域处理
	r.Use(Cors)
	r.Use(JWTAuth)

	r.Use(Casbin)
}

// 跳过验证
func LoginNoAuth(ctx iris.Context) (t bool) {
	path := ctx.Path()
	stringLogin := strings.SplitN(path, "/", -1)
	if stringLogin[len(stringLogin)-1] == "login" {
		//fmt.Printf("%v  login头--\n", stringLogin[len(stringLogin)-1] == "login")
		return true
	} else {
		return false
	}
}
