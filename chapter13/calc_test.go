package main

import (
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"
)

// ============ 1. 表驱动测试(table-driven test) + 子测试(subtest) ============
// Go 社区最主流的测试范式：把用例写成一张表，循环里对每行跑一次断言。
// 用 t.Run 给每个用例起名字，失败时能精确定位是哪个用例挂了。
func TestDivide(t *testing.T) {
	cases := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{"正常整除", 10, 2, 5, nil},
		{"向下取整", 7, 2, 3, nil},
		{"除以零", 1, 0, 0, ErrDivByZero},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { // 子测试：go test -run=TestDivide/除以零 可单独跑
			got, err := Divide(c.a, c.b)

			// 用 errors.Is 断言错误，而不是直接比 == 或比字符串
			if !errors.Is(err, c.wantErr) {
				t.Fatalf("Divide(%d,%d) 错误 = %v, 期望 %v", c.a, c.b, err, c.wantErr)
			}
			if err == nil && got != c.want {
				t.Errorf("Divide(%d,%d) = %d, 期望 %d", c.a, c.b, got, c.want)
			}
		})
	}
}

// ============ 2. 基准测试(benchmark) ============
// 函数名以 Benchmark 开头，跑 `go test -bench=. -benchmem .` 观察 ns/op、allocs/op。
func BenchmarkFib(b *testing.B) {
	for b.Loop() { // Go 1.24+ 写法，替代 for i := 0; i < b.N; i++
		_ = Fib(30)
	}
}

// ============ 3. 示例测试(example) ============
// 以 Example 开头，末尾的 // Output: 注释会被 go test 当断言执行；
// 同时这些示例会展示在 pkg.go.dev 文档里，一举两得。
func ExampleReverse() {
	fmt.Println(Reverse("hello"))
	fmt.Println(Reverse("你好"))
	// Output:
	// olleh
	// 好你
}

// ============ 4. 模糊测试(fuzzing, Go 1.18+) ============
// 以 Fuzz 开头，f.Add 提供种子语料，f.Fuzz 的函数会被喂入引擎自动生成的大量输入。
// 这里验证一个"性质"：反转两次应还原，且结果仍是合法 UTF-8。
// 运行：go test -fuzz=FuzzReverse .   （会一直跑，Ctrl+C 停止；失败输入自动存到 testdata/）
func FuzzReverse(f *testing.F) {
	for _, seed := range []string{"", "a", "hello", "你好", "Go语言"} {
		f.Add(seed) // 种子语料
	}
	f.Fuzz(func(t *testing.T, orig string) {
		// Reverse 是按 rune 反转的，只对「合法 UTF-8」保证性质成立。
		// 非法 UTF-8（如孤立的 "\xb8"）经 []rune 转换会被替换成 U+FFFD，无法还原，
		// 这不是 bug 而是约定之外的输入，直接跳过。fuzz 引擎很快就能构造出这类输入。
		if !utf8.ValidString(orig) {
			t.Skip("跳过非法 UTF-8 输入")
		}
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("反转两次未还原: orig=%q doubleRev=%q", orig, doubleRev)
		}
		// 合法 UTF-8 反转后应仍是合法 UTF-8（这正是为什么必须按 rune 反转）
		if !utf8.ValidString(rev) {
			t.Errorf("反转破坏了 UTF-8: orig=%q rev=%q", orig, rev)
		}
	})
}
