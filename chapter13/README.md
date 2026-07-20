# 测试

> 约定：测试文件名以 `_test.go` 结尾，和被测代码放同一个包；`go test` 会自动发现并编译它们。

## 四种测试函数

| 前缀 | 签名 | 作用 | 运行 |
| ---- | ---- | ---- | ---- |
| `Test` | `func TestXxx(t *testing.T)` | 单元测试 | `go test` |
| `Benchmark` | `func BenchmarkXxx(b *testing.B)` | 性能基准 | `go test -bench=.` |
| `Example` | `func ExampleXxx()` | 示例 + 文档 + 断言 | `go test`（对比 `// Output:`） |
| `Fuzz` | `func FuzzXxx(f *testing.F)` | 模糊测试 | `go test -fuzz=FuzzXxx` |

## 表驱动测试（主流范式）

> 把用例写成一张表，循环里逐行断言；用 `t.Run` 给每个用例命名。

```go
cases := []struct {
	name    string
	a, b    int
	want    int
	wantErr error
}{
	{"正常整除", 10, 2, 5, nil},
	{"除以零", 1, 0, 0, ErrDivByZero},
}
for _, c := range cases {
	t.Run(c.name, func(t *testing.T) { ... })
}
```

- 加用例只改表，测试逻辑不动，可读性和可维护性都好
- 子测试可单独跑：`go test -run=TestDivide/除以零`
- 断言错误用 `errors.Is`，不要比字符串（承接 chapter11）
- `t.Error` 记录失败但继续；`t.Fatal` 记录并立即中止当前测试

## 基准测试

```go
func BenchmarkFib(b *testing.B) {
	for b.Loop() { // Go 1.24+ 推荐，替代 for i := 0; i < b.N; i++
		_ = Fib(30)
	}
}
```

```bash
go test -bench=. -benchmem .   # 关注 ns/op（快慢）和 allocs/op（分配次数）
```

## 示例测试（示例即文档即断言）

```go
func ExampleReverse() {
	fmt.Println(Reverse("hello"))
	// Output:
	// olleh
}
```

- 末尾 `// Output:` 注释会被当断言：实际输出不符就测试失败
- 同时这些示例会展示在 `pkg.go.dev` 文档里，一举两得

## 模糊测试 (Fuzzing)

> Go 1.18+ 内置。只提供**种子语料**和一个要满足的**性质**，引擎自动生成海量刁钻输入去打破它。

```go
func FuzzReverse(f *testing.F) {
	f.Add("你好")                 // 种子
	f.Fuzz(func(t *testing.T, orig string) {
		if Reverse(Reverse(orig)) != orig { // 性质：反转两次应还原
			t.Errorf(...)
		}
	})
}
```

```bash
go test -fuzz=FuzzReverse .   # 持续运行，Ctrl+C 停止
```

- 找到反例后，失败输入会自动写入 `testdata/fuzz/`，之后作为普通用例长期回归
- 特别擅长揪出边界 bug：空串、非法 UTF-8、超长输入等（本例按 `rune` 反转就是为了通过 fuzz）

## 常用命令速查

```bash
go test ./...               # 跑当前模块所有测试
go test -v ./...            # 显示每个用例
go test -run=TestDivide .   # 按正则筛选测试
go test -race ./...         # 竞态检测（并发代码必跑，承接 chapter09）
go test -cover ./...        # 覆盖率百分比
go test -coverprofile=c.out . && go tool cover -html=c.out  # 可视化覆盖率
```

## 底层原理解析

- **`go test` 会生成一个临时的测试主程序**：把包内的 `_test.go` 编译进去，自动发现 `Test/Benchmark/Example/Fuzz` 函数并调度执行，跑完再清理
- **`b.N` / `b.Loop()` 的自适应**：基准框架会不断加大迭代次数，直到累计耗时足够稳定，再算出 `ns/op`。`b.Loop()`（Go 1.24+）还能阻止编译器把"结果没被使用"的被测代码优化掉
- **`-race` 的原理**：编译时插桩，运行时用 happens-before 关系检测对同一内存的无同步并发读写。它只能发现**实际执行到**的竞态，所以要配合有代表性的测试用例
- **黑盒测试包 `xxx_test`**：测试文件可声明为 `package foo_test`（外部测试包），只能访问导出标识符，用来验证"从使用者视角"的 API；本章用的是同包白盒测试（能访问未导出成员）

## 避坑指南

- **测试之间共享状态**：全局变量、临时文件、数据库若不隔离，用例顺序一变就飘。每个子测试自带干净数据，或用 `t.Cleanup()` 收尾
- **基准结果被优化掉**：被测代码结果没被使用时可能被整段删除，测出假的"0 ns/op"。用 `b.Loop()` 或把结果赋给包级 `sink` 变量
- **忽略 `-race`**：并发代码不跑 `-race` 等于没测。它是发现数据竞争几乎唯一可靠的手段
- **Example 的 `// Output:` 顺序敏感**：`map` 遍历无序，别在示例里直接打印 map 的多个键值，否则时对时错；需要就先排序或用 `// Unordered output:`
- **fuzz 语料入库要提交**：`testdata/fuzz/` 下的失败样例应纳入版本控制，它们是宝贵的回归用例
- **只追覆盖率数字**：100% 覆盖 ≠ 没 bug。覆盖率只说明"这行被执行过"，不代表"断言了正确行为"。优先覆盖核心逻辑和边界，而非凑百分比
