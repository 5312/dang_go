package controller

import (
	// go 标准包

	"dang_go/internal/database"

	"dang_go/internal/model/system"
	"fmt"

	// 第三方包
	"github.com/gogf/gf/util/gconv"
	"github.com/kataras/iris/v12"
)

type Curder interface {
	AddMenu(ctx iris.Context)
	DeleteMenu(ctx iris.Context)
	UpMenu(ctx iris.Context)
	GetListMenu(ctx iris.Context)
}

type Menus struct {
	Name string
}

/* 增 */
func (m *Menus) AddMenu(ctx iris.Context) {
	// 接收参数
	var creatMenu system.Menu
	if err := ctx.ReadJSON(&creatMenu); err != nil {

		result := database.DB.Create(&creatMenu) // 通过数据的指针来创建
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
	arr := []system.Menu{}
	arr = append(arr, creatMenu)
	ctx.JSON(
		iris.Map{
			"Success": true,
			"Code":    1,
			"Msg":     "添加失败",
			"data":    arr,
		})
}

/* 删 */
func (m *Menus) DeleteMenu(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("ID", 0)

	err := database.DB.Delete(&system.Menu{}, id).Error
	if err != nil {
		ctx.JSON(iris.Map{
			"success": false,
			"id":      id,
			"msg":     "删除失败",
		})
		return
	}
	ctx.JSON(iris.Map{
		"success": true,
		"id":      id,
		"msg":     "删除成功",
	})
}

/* 改 */
func (m *Menus) UpMenu(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("ID", 0)

	var results []*system.Menu //[]map[string]interface{}
	// 从 t 中获取 返回行数 RowsAfffected
	t := database.DB.Table("menus").First(&results, id).Error

	err := database.DB.Save(&results).Error

	if err != nil && t != nil {
		ctx.JSON(iris.Map{
			"success": false,
			"id":      id,
			"msg":     "删除失败",
		})
		return
	}
	ctx.JSON(iris.Map{
		"success": true,
		"id":      id,
		"msg":     "修改成功",
	})
}

/* 查 */
func (m *Menus) GetListMenu(ctx iris.Context) {
	// Get all records
	var results []*system.Menu //[]map[string]interface{}
	// 从 t 中获取 返回行数 RowsAfffected
	t := database.DB.Table("menus").Find(&results)
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

/* 返回处理 */
func backJson(successOrError bool) interface{} {

	if successOrError {
		res := Res{
			Success: true,
			Code:    0,
			Msg:     "添加成功",
		}
		return res
	} else {
		res := Res{
			Success: true,
			Code:    1,
			Msg:     "添加失败",
		}
		return res
	}

}

/* 转换树结构 */
func TreeNode(list []*system.Menu, pid uint) []TreeResponse {
	var array []TreeResponse

	for _, item := range list {
		// pid == 0 是第一层
		if item.ParentId == pid {
			var children TreeResponse

			err := gconv.Struct(item, &children)
			if err != nil {
				panic(err)
			}
			children.Children = TreeNode(list, item.Model.ID)
			array = append(array, children)
		}
	}
	return array
}
