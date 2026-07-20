package main

import "fmt"

// 逃逸分析(escape analysis)：编译器在**编译期**决定一个变量分配在栈上还是堆上。
//   - 栈：随函数返回自动回收，零 GC 压力，快
//   - 堆：需要 GC 追踪回收，有开销
// 一句话规则：编译器只要**无法证明**变量的生命周期不超过当前函数，就把它挪到堆上（逃逸）。
//
// 观察方法（务必亲自跑）：
//	go build -gcflags="-m" .
// 会打印类似：
//	./escape.go:NN:M: moved to heap: u
//	./escape.go:NN:M: &User{...} escapes to heap

type User struct {
	Name string
	Age  int
}

// newUserEscape 返回了局部变量的地址 → 该变量在函数返回后仍被引用 → 必须逃逸到堆。
func newUserEscape(name string) *User {
	u := User{Name: name, Age: 18} // -gcflags="-m" 会提示 moved to heap: u
	return &u
}

// sumNoEscape 里的 s 只在函数内部使用，不会逃逸，稳稳待在栈上。
func sumNoEscape(nums []int) int {
	s := 0 // 不逃逸
	for _, n := range nums {
		s += n
	}
	return s
}

func escapeDemo() {
	u := newUserEscape("alice")
	fmt.Printf("newUserEscape 返回堆对象: %+v\n", *u)

	fmt.Println("sumNoEscape([1,2,3]) =", sumNoEscape([]int{1, 2, 3}))

	// 常见"隐性逃逸"来源，运行时看不出来，但 -gcflags="-m" 会告诉你：
	//   1. 把具体值赋给接口(interface{}/any) —— 装箱通常逃逸
	//   2. 闭包捕获的变量
	//   3. 变量太大、大小在编译期不确定（如 make 的长度是变量）
	//   4. 被 fmt.Println 这类 ...any 参数接收（会转成接口）
	var boxed any = 42 // 这次装箱就是一次典型逃逸
	fmt.Println("装箱到接口会逃逸:", boxed)
}
