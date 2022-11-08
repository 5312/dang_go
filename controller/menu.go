package controller

import (
	// go 标准包

	// 内部包
	"com.example.dang/internal/model"
	"com.example.dang/internal/model/system"

	// 第三方包
	"github.com/kataras/iris/v12"
)

// 错误处理方案： 错误类型
type ErrorShowType int

const (
	SILENT        ErrorShowType = 0 // silent
	WARN_MESSAGE                = 1
	ERROR_MESSAGE               = 2
	NOTIFICATION                = 3
	REDIRECT                    = 9 // redirect
)

type Response struct {
	Success      bool          `json:"success"`
	Total        int64         `json:"total"`
	Data         []system.Menu `json:"data"`
	ErrorCode    int           `json:"errorCode"`
	ErrorMessage string        `json:"errorMessage"`
	ShowType     ErrorShowType `json:"showType"`
	Page         int           `json:"page"`
}

func Menu(ctx iris.Context) {
	// Get all records
	var results []system.Menu //[]map[string]interface{}
	t := model.DB.Table("menus").Find(&results)

	res := Response{
		Success: true,
		Total:   t.RowsAffected,
		Data:    results,
	}
	ctx.JSON(res)

}
