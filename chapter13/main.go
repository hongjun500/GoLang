package main

import "fmt"

// 本章的重点全在 _test.go 文件里，主程序只是被测对象的载体。
// 请直接运行：
//	go test -v ./...              // 跑所有测试
//	go test -run=TestDivide ./... // 只跑某个测试
//	go test -bench=. ./...        // 跑基准测试
//	go test -race ./...           // 竞态检测
//	go test -cover ./...          // 覆盖率
//	go test -fuzz=FuzzReverse .   // 模糊测试（Ctrl+C 停止）
func main() {
	fmt.Println("被测对象见 calc.go / strutil.go；测试见 *_test.go。请运行 go test -v ./...")
}
