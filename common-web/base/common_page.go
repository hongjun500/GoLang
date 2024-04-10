// @author hongjun500
// @date 2024-04-10-21:02
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: 通用分页数据
// 实现分页接口： Paginator

package base

type CommonPage struct {
	PageNum   int    `json:"page_num"`   // 当前页
	PageSize  int    `json:"page_size"`  // 每页数量
	TotalPage int64  `json:"total_page"` // 总页数
	Total     int64  `json:"total"`      // 总记录数
	PageData  any    `json:"page_data"`  // 分页数据
	OrderBy   string `json:"order_by"`   // 排序
	Sort      string `json:"sort"`       // 排序方式
}

// Paginator 分页器接口
type Paginator interface {
	Paginate() error
}

// DatabaseHandler 数据库处理器接口
type DatabaseHandler interface {
	ExecuteCountSQL(SQL string, args ...any) (int64, error)
	ExecuteQuerySQL(dest any, SQL string, args ...any) error
}
