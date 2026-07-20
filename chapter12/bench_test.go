package main

import (
	"strings"
	"testing"
)

// 一个经典对比：字符串拼接的两种写法，性能相差一到两个数量级。
// 跑 `go test -bench=. -benchmem .` 亲眼看 ns/op 和 allocs/op 的差距。
// 采 profile 时请把包路径放前面，例如：
//   go test . -bench=BenchmarkConcat -cpuprofile=cpu.out
// （PowerShell 下 `-cpuprofile=cpu.out .` 会被拆成两个包而报错。）

const n = 1000

// concatPlus 用 += 拼接：每次都生成新字符串、复制一遍，O(n^2) 的拷贝 + 海量分配。
func concatPlus() string {
	s := ""
	for i := 0; i < n; i++ {
		s += "x"
	}
	return s
}

// concatBuilder 用 strings.Builder：内部是可增长的 []byte，摊还 O(n)，几乎零多余分配。
func concatBuilder() string {
	var b strings.Builder
	b.Grow(n) // 预分配容量，进一步减少扩容
	for i := 0; i < n; i++ {
		b.WriteByte('x')
	}
	return b.String()
}

// BenchmarkConcatPlus 会暴露大量分配：-benchmem 下 allocs/op 很高。
func BenchmarkConcatPlus(b *testing.B) {
	for b.Loop() { // Go 1.24+ 推荐写法，替代手写 for i := 0; i < b.N; i++
		_ = concatPlus()
	}
}

// BenchmarkConcatBuilder 应明显更快、分配更少。
func BenchmarkConcatBuilder(b *testing.B) {
	for b.Loop() {
		_ = concatBuilder()
	}
}
