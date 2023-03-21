package middleware

import (
	"github.com/kataras/iris/v12"
)

// Cors 实现服务端跨域
// Cors is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Cors(c iris.Context) {
	if c.Request().Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		//c.AbortWithStatus(200)
		c.StatusCode(204)
		c.Next()
	}
}
