package middleware

import (
	"dang_go/internal/enforcer"
	"dang_go/tools/app"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/x/errors"
)

// EnforcerTool
func Casbin(c iris.Context) {

	if LoginNoAuth(c) {
		c.Next()
		return
	}

	// 获取jwt中间件传递的信息
	claims := c.Values().Get("claims")
	fmt.Printf("%v 用户信息-- \n", claims)

	e := enforcer.EnforcerTool()

	//获取用户的角色
	sub := "admin"
	//获取请求的URI
	obj := c.Path()
	//获取请求方法
	act := c.Method()

	//判断策略中是否存在
	if ok, _ := e.Enforce(sub, obj, act); ok {
		fmt.Println("恭喜您,权限验证通过")
		c.Next()
	} else {
		fmt.Println("很遗憾,权限验证没有通过")

		c.StatusCode(401)
		error := errors.New("很遗憾,权限验证没有通过")
		app.Error(c, -1, error, "")
		c.StopExecution() //中止响应
		//c.Next()
	}
}
