package controller

import (
	// go 标准包

	// 内部包
	"dang_go/internal/model"
	"dang_go/internal/model/system"
	"fmt"

	// 第三方包
	"github.com/kataras/iris/v12"
)

type CurdMenu interface {
	GetListMenu()
	DeleteMenu()
	AddMenu()
}

type Menu struct{}

func (m *Menu) AddMenu(ctx iris.Context) {
	path := ctx.Path()
	fmt.Println("add")

	fmt.Println(path)
	// ctx.JSON()

}

func (m *Menu) GetListMenu(ctx iris.Context) {
	// Get all records
	var results []system.Menu //[]map[string]interface{}
	t := model.DB.Table("menus").Find(&results)

	res := Response{
		Success: true,
		Data:    results,
		Code:    0,
		Msg:     "请求成功",
		TablePage: &TablePage{
			Total: t.RowsAffected,
		},
	}
	ctx.JSON(res)
}

func (m *Menu) DeleteMenu(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("ID", 0)

	err := model.DB.Delete(&system.Menu{}, id).Error
	if err != nil {
		ctx.JSON(iris.Map{
			"success": false,
			"id":      id,
		})
		return
	}
	ctx.JSON(iris.Map{
		"success": true,
		"id":      id,
	})
}
