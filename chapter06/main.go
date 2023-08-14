package main

import (
	"fmt"
	"reflect"
	"runtime"
)

const GO_LANG = "GOlANG"

func main() {
	Print("hello go")

	fmt.Println(sum(1, 1))

	// 有参无返回值
	logNoReturn("hello, go")

	// 无参且无返回值
	logNoReturnAndParam()

	// 函数可以有多个参数并且有多个返回值
	arr1, arr2, arr3 := moreReturn(1, 2, 3)
	/*intArrays := make([]int, 3)
	stringArrays := make([]string, 1)
	intArrays[0] = arr1
	intArrays[1] = arr2
	stringArrays[0] = arr3
	fmt.Println(intArrays, stringArrays)*/

	// 打印时可接受参数%v
	fmt.Printf("返回的三个参数分别是%v,%v,%v \n", arr1, arr2, arr3)

	// 匿名函数  (隐藏函数名)
	var fc = func() int {
		i := 0
		i++
		fmt.Println("匿名函数作为变量使用")
		return i
	}
	fmt.Println(fc())

	func() {
		// func(){}()
		fmt.Println("直接使用匿名函数")
	}() /*这个小括号不能省略*/

	str := "外部变量str"
	func() {
		fmt.Printf("直接使用匿名函数引用外部变量=%v\n", str)
		fmt.Printf("直接使用匿名函数引用外部常量=%v\n", GO_LANG)
	}()

	// defer关键字配合闭包使用

	strWord := "hello world"
	defer func() { // 这里的defer关键字会使此匿名函数延迟调用, 并且strWord的内存地址会被defer关键字保留，而并不是保留其值的内容
		fmt.Printf("外部变量strWord=%v \n", strWord)
	}()

	strWord = "hello goLang"             // 这里改变了变量strWord的值，内存地址并没有改变
	fmt.Printf("strWord=%v \n", strWord) // 这里执行完之后匿名函数开始调用，记录的strWord的值就是改变之后的值

	// 函数式编程 (打印出main.div)
	fmt.Println(func_div(div, 100, 2))
	// 匿名函数， 打印出main.main.func(随机数命名值)
	fmt.Println(func_div(func(i int, i2 int) (int, int) {
		return i / i2, i % i2
	}, 100, 2))

	argsFunc(1)
	argsFunc(1, 2, 3, -9999)

}

// 函数关键字func 函数名pring, 入参为一个string类型，无返回值
func Print(str string) {
	fmt.Printf("传入了一个参数：%v \n", str)
	fmt.Println(str)
}

/*
一个string类型入参并打印，无法返回值
*/
func logNoReturn(str string) {
	fmt.Println(str)
}

/*
两个int类型入参，返回一个int类型参数
*/
func sum(x, y int) int {
	return x + y
}

/*
无参数无返回值的打印
*/
func logNoReturnAndParam() {
	fmt.Println("无参数无返回值的打印")
}

/*
接收三个参数，并且返回两个int类型一个string类型的值
注意这里的返回值a,b,c被命名了
返回值的命名应当具有一定含义，可以作为文档使用
返回值命名时，返回最后可以直接return
*/
func moreReturn(x, y, z int) (a, b int, c string) {
	a = x + y + z
	b = x * y * z
	// return a, b, "string"
	// 效果同上return,如果函数比较长这种方式可读性比较差
	return
}

// 对比上一个函数
func moreReturnV1(x, y, z int) (int, int, string) {
	// := 语句可以省略var声明
	// 但是 := 只能在函数内声明，函数外只能以关键字开头
	a := x + y + z
	b := x * y * z
	return a, b, "string"
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func func_div(fc func(int, int) (int, int), c, d int) (int, int) {
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
	fmt.Println("argsFunc函数调用")
	for _, arg := range args {
		fmt.Println(arg)
	}
}
