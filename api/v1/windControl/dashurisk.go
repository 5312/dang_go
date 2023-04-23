package windControl

import (
	"dang_go/controller/risk"
	"github.com/kataras/iris/v12/core/router"
)

// RegisterRisk 天狼星 大数风控
func RegisterRisk(v1 router.Party) {
	// risk 组
	api := v1.Party("/risk")
	{
		// 天狼星
		api.Post("/getTestDBInfo", risk.TestDashu)
		api.Post("/getDBInfo", risk.Dashu)

	}
}
