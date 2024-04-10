// @author hongjun500
// @date 2024-04-10-22:00
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: rdb_handler.go

package base

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseHandlerCreator func(dataSourceName string) (DatabaseHandler, error)

var NewMySQLHandler DatabaseHandlerCreator = func(dataSourceName string) (DatabaseHandler, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return &MySQLHandler{DB: db}, nil
}

var NewOracleHandler DatabaseHandlerCreator = func(dataSourceName string) (DatabaseHandler, error) {
	db, err := sql.Open("oracle", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return &OracleHandler{DB: db}, nil
}

// MySQLHandler MySQL 数据库处理器
type MySQLHandler struct {
	DB *sql.DB
}

// OracleHandler Oracle 数据库处理器
type OracleHandler struct {
	DB *sql.DB
}

// ExecuteCountSQL 执行查询总数 sql
func (handler *MySQLHandler) ExecuteCountSQL(SQL string, args ...any) (int64, error) {
	return executeSQL(handler.DB, SQL, args...)
}

// ExecuteQuerySQL 执行查询数据 sql
func (handler *MySQLHandler) ExecuteQuerySQL(dest any, SQL string, args ...any) error {
	return executeQuerySQL(handler.DB, dest, SQL, args...)
}

func (handler *OracleHandler) ExecuteCountSQL(SQL string, args ...any) (int64, error) {
	return executeSQL(handler.DB, SQL, args...)
}

func (handler *OracleHandler) ExecuteQuerySQL(dest any, SQL string, args ...any) error {
	return executeQuerySQL(handler.DB, dest, SQL, args...)
}

func executeSQL(db *sql.DB, SQL string, args ...any) (int64, error) {
	var count int64
	err := db.QueryRow(SQL, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to execute count SQL: %v", err)
	}
	return count, nil
}

func executeQuerySQL(db *sql.DB, dest any, SQL string, args ...any) error {
	rows, err := db.Query(SQL, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query SQL: %v", err)
	}
	defer rows.Close()

	destValue := reflect.ValueOf(dest)
	destElemType := destValue.Elem().Type().Elem()

	for rows.Next() {
		row := reflect.New(destElemType).Elem()
		scanArgs := make([]any, row.NumField())
		for i := 0; i < row.NumField(); i++ {
			scanArgs[i] = row.Field(i).Addr().Interface()
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}

		destValue.Elem().Set(reflect.Append(destValue.Elem(), row.Addr().Elem()))
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("failed to iterate rows: %v", err)
	}

	return nil
}
