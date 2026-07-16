package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const GoLang = "GoLang"

func main() {

	fmt.Println(sum(1, 1))

	// 有参无返回值
	logNoReturn("hello, go")

	// 无参且无返回值
	logNoReturnAndParam()

	// 多返回值 + 命名返回值
	a, b, c := moreReturn(1, 2, 3)
	fmt.Printf("moreReturn 返回值: a=%v, b=%v, c=%v\n", a, b, c) // 可接受参数 	%v

	// 匿名函数  (隐藏函数名) 赋值给变量后调用
	var double = func(param int) int {
		fmt.Printf("匿名函数 double(param int) int 接收到的参数param = %d\n", param)
		return param * 2
	}
	fmt.Printf("double(100) = %v\n", double(100))

	// 匿名函数直接调用
	str := "外部变量 str "
	func() string {
		fmt.Printf("IIFE 捕获外部变量=%q, 外部常量=%q\n", str, GoLang)
		return str
	}() // 这个小括号不能省略，代表立即调用此匿名函数

	// 闭包：捕获的是"变量本身"，多次调用共享同一个 count（会逃逸到堆）
	counter := newCounter()
	fmt.Printf("counter(): %d %d %d\n", counter(), counter(), counter()) // 1 2 3

	// defer + 闭包 vs defer 立即求值
	deferValueVsClosure()

	// defer 的执行顺序：后进先出(LIFO)
	deferOrderDemo()

	// defer 修改命名返回值
	fmt.Println("withNamedReturn(2) =", withNamedReturn(2))

	// panic / recover：Go 处理运行时"异常"的惯用法
	safeDivide(10, 0) // 被 recover 捕获，程序不会崩溃
	safeDivide(10, 2)

	// 函数式编程 (打印出main.applyDiv) 把"函数"当作参数传入
	fmt.Println(applyDiv(div, 100, 2))
	// 匿名函数， 打印出main.main.func(编译器按匿名函数在源码中出现的顺序生成，形如 main.main.func1、main.main.func2……)
	fmt.Println(applyDiv(func(x, y int) (int, int) {
		return x / y, x % y
	}, 100, 2))

	// 可变参数(variadic)：本质是切片
	fmt.Println("sumAll() =", sumAll())
	fmt.Println("sumAll(1,2,3) =", sumAll(1, 2, 3))
	nums := []int{4, 5, 6}
	fmt.Println("sumAll(nums...) =", sumAll(nums...)) // 用 ... 展开切片

	// 可变参数陷阱：形参与调用方切片共享底层数组
	variadicPitfall()
}

// 有参无返回值
func logNoReturn(str string) {
	fmt.Println(str)
}

// 两个 int 类型的参数并且返回一个 int 类型的值
func sum(x, y int) int {
	return x + y
}

// 无参无返回值
func logNoReturnAndParam() {
	fmt.Println("无参数无返回值的函数")
}

// moreReturn 接收三个参数，并且返回两个 int 类型一个 string 类型的值
// 这里的返回值 a,b,c 被命名, 可以直接 return
// 返回值的命名应当具有一定含义，可以作为文档使用
func moreReturn(x, y, z int) (a, b int, c string) {
	a = x + y + z
	b = x * y * z
	// return a, b, "string"
	// 效果同上,如果函数比较长这种方式可读性比较差
	return
}

// 对比上一个函数
func moreReturnV1(x, y, z int) (int, int, string) {
	a := x + y + z
	b := x * y * z
	return a, b, "string"
}

// newCounter 返回一个闭包。每次调用返回的函数，count 都会 +1，
// 说明闭包捕获的是变量 count 本身（逃逸到堆），而非它的初始值快照。
func newCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// deferValueVsClosure 对比 defer 的两种求值时机：
//   - defer 直接调用函数：实参在 defer 语句处"立即求值"并拷贝
//   - defer + 闭包：闭包读取变量本身，拿到的是 return 时刻的最新值
func deferValueVsClosure() {
	str := "hello world"
	// 闭包：读到修改后的值 hello goLang
	defer func() {
		fmt.Printf("[defer 闭包] str = %q\n", str)
	}()
	// 立即求值：实参此刻即被拷贝，永远打印 hello world
	defer fmt.Printf("[defer 立即求值] str = %q\n", str)
	str = "hello goLang"
}

// deferOrderDemo 展示多个 defer 的执行顺序：后进先出(LIFO)，输出 3 2 1。
func deferOrderDemo() {
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("[defer 顺序] i = %d\n", i)
	}
}

// withNamedReturn 展示 defer 可以修改命名返回值：
// return 先把 result 置为 入参 param，defer 再把它改成 * 10，最终返回 param 10的倍数。
func withNamedReturn(param int) (result int) {
	defer func() {
		result = param * 10
	}()
	result = param
	return
}

// safeDivide 展示 panic/recover：recover 只有在 defer 中才生效，
// 用来把运行时 panic 转成可控处理，避免整个程序崩溃。
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[recover] 捕获到 panic: %v\n", r)
		}
	}()
	fmt.Printf("safeDivide(%d, %d) = %d\n", a, b, a/b) // b==0 时触发 panic
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

// applyDiv 是高阶函数：把"函数"当作参数传入并调用。
// 借助 reflect + runtime 打印出传入函数的名字，便于观察。
func applyDiv(fn func(int, int) (int, int), a, b int) (int, int) {
	full := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	// 具名函数形如 main.div；匿名函数形如 main.main.func2，取最后一段更直观
	name := full[strings.LastIndex(full, ".")+1:]
	fmt.Printf("调用函数 %s(%d, %d)\n", name, a, b)
	return fn(a, b)
}

// sumAll 可变参数：本质上 nums 就是一个 []int 切片；不传参时它是 nil。e.g. sumAll(1,2), sumAll(1,2,3,4,5)
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// variadicPitfall 演示可变参数经典陷阱：
// f(slice...) 会把调用方切片的底层数组"借"给形参，元素写操作会外泄。
func variadicPitfall() {
	data := []int{1, 2, 3}
	writeTail(data...) // ... 展开：形参与 data 共享同一底层数组
	fmt.Printf("[variadic 陷阱] 调用方 data = %v\n", data) // [1 2 -1]
}
func writeTail(nums ...int) {
	if len(nums) > 0 {
		nums[len(nums)-1] = -1 // 直接改元素，会影响调用方
	}
}
