package middleware

import "github.com/kataras/iris/v12"

func InitMiddleware(r *iris.Application) {
	// 跨域处理
	r.Use(Cors)
}
