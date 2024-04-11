// @author hongjun500
// @date 2024-04-10-22:15
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: rdb_handler_test.go

package base

import (
	"testing"

	"github.com/hongjun500/GoLang-master/common-web/base"
)

type Handler struct {
	Id      int64
	Handler string
}

var handlers = &[]Handler{}

func TestMySQLHandler(t *testing.T) {
	handler, err := base.NewMySQLHandler("root:hongjun500@tcp(127.0.0.1:3306)/common-web?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	if err != nil {
		t.Fatalf("failed to create MySQLHandler: %v", err)
	}

	count, err := handler.ExecuteCountSQL("SELECT COUNT(*) FROM database_handler where id = ? ", 1, 10, 0)
	if err != nil {
		t.Fatalf("failed to execute count SQL: %v", err)
	}

	t.Logf("count: %d", count)

	err = handler.ExecuteQuerySQL(handlers, "SELECT * FROM database_handler where id = ?", 1, 10, 0)
	if err != nil {
		t.Fatalf("failed to execute query SQL: %v", err)
	}

	t.Logf("handlers: %+v", handlers)
}

func TestOracleHandler(t *testing.T) {
	handler, err := base.NewOracleHandler("your_oracle_data_source_name")
	if err != nil {
		t.Fatalf("failed to create OracleHandler: %v", err)
	}

	count, err := handler.ExecuteCountSQL("SELECT COUNT(*) FROM your_table")
	if err != nil {
		t.Fatalf("failed to execute count SQL: %v", err)
	}

	t.Logf("count: %d", count)

	err = handler.ExecuteQuerySQL(handlers, "SELECT * FROM your_table")
	if err != nil {
		t.Fatalf("failed to execute query SQL: %v", err)
	}

	t.Logf("handlers: %+v", handlers)
}
