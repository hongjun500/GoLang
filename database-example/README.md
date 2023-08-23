# 访问关系型数据库 MySQL

## 依赖

- `database/sql`
- `github.com/go-sql-driver/mysql`

## 连接配置

~~~go
// 设置时区为东八区
loc, _ := time.LoadLocation("Asia/Shanghai")
// 数据库配置
config := mysql.Config{
User:   "root",
Passwd: "**",
Net:    "tcp",
Addr:   "localhost:3306",
DBName: "recordings",
Loc:    loc,
// 允许本地密码认证
AllowNativePasswords: true,
}
~~~

## 声明一个数据库句柄

`var db *sql.DB`

### 读操作

- 使用 `db.QueryRow` 查询单行数据
- 使用 `db.Query` 查询多行数据

### 写操作

- 使用 `db.Exec` 执行 `insert\delete\update` 语句