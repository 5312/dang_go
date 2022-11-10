package controller

import (
	// go 标准包

	// 内部包
	"dang_go/internal/model"
	"dang_go/internal/model/system"
	"fmt"

	// 第三方包
	"github.com/gogf/gf/util/gconv"
	"github.com/kataras/iris/v12"
)

type CurdMenu interface {
	GetListMenu()
	DeleteMenu()
	AddMenu()
}

type Menu struct{}

func (m *Menu) AddMenu(ctx iris.Context) {
	// 接收参数
	var creatMenu system.Menu
	if err := ctx.ReadJSON(&creatMenu); err != nil {
		// fmt.Println(creatMenu)
		result := model.DB.Create(&creatMenu) // 通过数据的指针来创建
		fmt.Printf("%v\n", result.Error)
		if result.Error == nil {
			ctx.JSON(Response{
				Res: &Res{
					Success: true,
					Code:    0,
					Msg:     "添加成功",
				},
			})
			return
		}

	}
	ctx.JSON(ResponseError{
		Res: &Res{
			Success: false,
			Code:    1,
			Msg:     "添加失败",
		},
	})

}

func (m *Menu) GetListMenu(ctx iris.Context) {
	// Get all records
	var results []*system.Menu //[]map[string]interface{}
	// 从 t 中获取 返回行数 RowsAfffected
	t := model.DB.Table("menus").Find(&results)
	a := TreeNode(results, 0)

	fmt.Println(a)

	res := Response{
		Res: &Res{Success: true,
			Code: 0,
			Msg:  "请求成功",
		},
		Data: a,
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

/* 转换树结构 */
func TreeNode(list []*system.Menu, pid uint) []TreeResponse {
	var array []TreeResponse

	for _, item := range list {
		// pid == 0 是第一层
		if item.Parent_id == pid {
			var children TreeResponse

			err := gconv.Struct(item, &children)
			if err != nil {
				panic(err)
			}
			children.Children = TreeNode(list, item.ID)
			array = append(array, children)
		}
	}
	return array
}
