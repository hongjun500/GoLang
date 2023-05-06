# 访问关系型数据库 MySQL
## 依赖
- `database/sql`
- `github.com/go-sql-driver/mysql`

## 声明一个数据库句柄
`var db *sql.DB`

### 读操作
- 使用 `db.QueryRow` 查询单行数据
- 使用 `db.Query` 查询多行数据

### 写操作
- 使用 `db.Exec` 执行 `insert\delete\update` 语句