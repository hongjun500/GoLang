package main

import (
	"fmt"
	"reflect"
	"runtime"
)

const GO_LANG = "GOlANG"

func main() {

	fmt.Println(sum(1, 1))

	// 有参无返回值
	logNoReturn("hello, go")

	// 无参且无返回值
	logNoReturnAndParam()

	// 函数可以有多个参数并且有多个返回值
	a, b, c := moreReturn(1, 2, 3)

	// 打印时可接受参数 %v
	fmt.Printf("返回的三个参数分别是%v,%v,%v \n", a, b, c)

	// 匿名函数  (隐藏函数名)
	var fc = func(name string) int {
		i := 0
		i++
		fmt.Printf("匿名函数接收到的参数是：%s\n", name)
		return i
	}
	fc("fc")

	str := "外部变量 str "
	func() string {
		fmt.Printf("直接使用匿名函数时引用外部变量=%v\n", str)
		fmt.Printf("直接使用匿名函数时引用外部常量=%v\n", GO_LANG)
		return str
	}() // 这个小括号不能省略，代表调用此匿名函数

	// defer 关键字配合闭包使用
	strWord := "hello world"
	defer func() { // 这里的 defer 关键字会使此匿名函数延迟调用, 并且 strWord 的内存地址会被匿名函数保留，而并不是保留其值
		fmt.Printf("外部变量 strWord =%v \n", strWord) // hello goLang
	}()

	// 这里改变了变量 strWord 的值，内存地址并没有改变
	strWord = "hello goLang"

	// 函数式编程 (打印出main.div)
	fmt.Println(funcDiv(div, 100, 2))
	// 匿名函数， 打印出main.main.func(随机数命名值)
	fmt.Println(funcDiv(func(i int, i2 int) (int, int) {
		return i / i2, i % i2
	}, 100, 2))

	argsFunc(1)
	argsFunc(1, 2, 3, -9999)

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

func div(a, b int) (int, int) {
	return a / b, a % b
}

func funcDiv(fc func(int, int) (int, int), c, d int) (int, int) {
	pointer := reflect.ValueOf(fc).Pointer()
	funcName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("funcation with %s \n args=%d, \n args=%d \n", funcName, c, d)
	return fc(c, d)
}

// 不定项参数可以传入多个参数并且类型都是数值类型
// argsFunc(1,2)
// argsFunc(1)
func argsFunc(args ...int) {
	// 得到一个新的切片
	args = append(args, 1888888)
	fmt.Println("argsFunc 函数调用")
	for _, arg := range args {
		fmt.Println(arg)
	}
}
