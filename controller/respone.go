package controller

import "dang_go/internal/model/system"

type TablePage struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

// 必返字段
type Res struct {
	Success bool   `json:"success"`
	Code    int    `json:"code" t:"0成功1失败"`
	Msg     string `json:"msg"`
}

// 成功
type Response struct {
	*Res
	Data []TreeResponse `json:"data"`
	// Table   *TablePage    `json:"table"`
	*TablePage
}

// 树结构
type TreeResponse struct {
	*system.Menu
	Children []TreeResponse `json:"children"`
}

// 失败
type ResponseError struct {
	*Res
}
