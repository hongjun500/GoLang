// @author hongjun500
// @date 2024-04-10-21:20
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: 原始 sql 分页

package base

type NativeSqlPage struct {
	DB                 DatabaseHandler // 数据库处理器 例如: mysql、postgres、sqlite3、sqlserver、oracle
	CountSQL, QuerySQL string          // 查询总数 sql 例如: select count(*) from table where id = ?, 查询数据 sql 例如: select * from table where id = ?
	QueryArgs          []any           // 查询数据参数
	*CommonPage                        // 通用分页数据
}

func NewNativeSqlPage(db DatabaseHandler, countSQL, querySQL string, queryArgs []any) *NativeSqlPage {
	return &NativeSqlPage{
		DB:         db,
		CountSQL:   countSQL,
		QuerySQL:   querySQL,
		QueryArgs:  queryArgs,
		CommonPage: &CommonPage{},
	}
}

func (page *NativeSqlPage) Paginate() error {
	count, err := page.DB.ExecuteCountSQL(page.CountSQL, page.QueryArgs...)
	if err != nil {
		return err
	}
	page.Total = count
	page.TotalPage = count / int64(page.PageSize)
	if count%int64(page.PageSize) > 0 {
		page.TotalPage++
	}
	err = page.DB.ExecuteQuerySQL(&page.PageData, page.QuerySQL, page.QueryArgs...)
	if err != nil {
		return err
	}
	return nil
}
