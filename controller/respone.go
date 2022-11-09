package controller

import "dang_go/internal/model/system"

type TablePage struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}
type Response struct {
	Success bool          `json:"success"`
	Code    int           `json:"code"`
	Msg     string        `json:"msg"`
	Data    []system.Menu `json:"data"`
	// Table   *TablePage    `json:"table"`
	*TablePage
}
