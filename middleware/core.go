package middleware

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

// Cors 实现服务端跨域
func Cors(ctx iris.Context) {
	fmt.Printf("跨域设置")
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "*")
		// ctx.Header("Access-Control-Allow-Headers", "Content-Type, Api, Accept, Authorization, Version, Token")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
