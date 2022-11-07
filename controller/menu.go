package controller

import (
	"fmt"

	"com.example.dang/internal/model"
	"com.example.dang/internal/model/system"
	"github.com/kataras/iris/v12"
)

/*
// 与后端约定的响应数据格式

	interface ResponseStructure {
	  success: boolean;
	  data: any;
	  errorCode?: number;
	  errorMessage?: string;
	  showType?: ErrorShowType;
	}
*/
// 错误处理方案： 错误类型
type n int

const (
	SILENT        n = 0 // silent
	WARN_MESSAGE    = 1
	ERROR_MESSAGE   = 2
	NOTIFICATION    = 3
	REDIRECT        = 9 // redirect
)

type Response struct {
	success      bool
	total        int64
	data         []system.Menu
	errorCode    int
	errorMessage string
	showType     n
}

func Menu(ctx iris.Context) {
	// user := system.Menu{Name: "Jinzhu"}

	// model.DB.Create(&user)

	// Get all records
	// result := model.DB.Find(&menu)
	var results []system.Menu //[]map[string]interface{}
	t := model.DB.Table("menus").Find(&results)
	//db.Model(&User{})
	// fmt.Printf("user:%#v\n", results)
	// fmt.Printf("user:%#v\n", t)
	/* iris.Map{"message": "Hello Iris!"} */
	res := Response{
		success: true,
		total:   t.RowsAffected,
		data:    results,
	}
	fmt.Printf("user:%#v\n", res)
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	// ctx.View("index.html")

}
