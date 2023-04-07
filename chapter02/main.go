package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	// 循环只有for
	// 不需要小括号，但大括号是必须
	// 分号;隔开的三个表达式
	// 初始化语句，条件表达式，后置语句
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	fmt.Println("sum值", sum)

	// 初始化语句和后置语句不是必须 例
	i := 100 // 本质上这里的i赋值100就是初始化语句
	for i > 0 {
		i-- // 而这里的i--则是后置语句，作用其实就是改变初始语句赋的值
		// 使其条件表达式的值为false结束循环，否则将陷入无限循环
	}
	fmt.Println("i值", i)

	fmt.Println("break关键字分割线-------------------------------------")

	// break关键字可以结束其所在循环结构体的执行（这个示例会打印出4个0123456789）
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
			// 如果满足了i==0时就会结束所在的一层循环
		}
		for i := 0; i < 10; i++ {
			fmt.Print(i)
		}
		fmt.Println("\n")
	}
	fmt.Println("continue关键字分割线-------------------------------------")

	// continue关键字可以结束条件判断其所在循环的执行（这个示例会打印出10个012356789）
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			if i == 4 {
				continue
				// 如果满足了i==0时就会跳出这个满足这个条件的循环，而不是循环体
			}
			fmt.Print(i)
		}
		fmt.Println("\n")
	}
	fmt.Println("分割线-------------------------------------")

	sum = 0
	for i < 100 {
		sum = sum + i
		i++
	}
	fmt.Println("更改后i值", sum)

	// 下例类似于java中的while
	a := 0
	for a < 10 {
		// if语句不需要小括号，必需大括号，类似for
		if a != 0 && a%2 == 0 {
			fmt.Println("偶数", a)
		}
		a++
	}

	value := pow(2, 3, 100)
	fmt.Println(value)

	/*fmt.Println(Sqrt(2))
	fmt.Println(SqrtFloat64(2))*/

	// 这两个结果不一样
	fmt.Println(1 / 2)
	fmt.Println(1 / 2.0)

	// switch语句 从上向下执行(可以没有小括号，大括号必须)
	// 同if一样可以有短变量声明
	// 每个case结构体后有默认的break,
	// 并且只会运行表达式对应选中的case，此时switch运行到此结束
	// 如果选中了case并且对应case包含一个fallthrough关键字语句，那么紧随其后的case也会执行

	// 可以没有default语句

	// 这个switch只会打印出Go
	switch lang := "Go"; lang {
	case "Go":
		fmt.Println("Go.")
	case "Java":
		fmt.Println("Java.")
		fallthrough
	case "Python":
		fmt.Println("Python.")

		/*default:
		fmt.Printf("%s.\n", lang)*/
	}

	// 这个switch会打印出Go And Java
	switch lang2 := "Go"; lang2 {
	case "Go":
		fmt.Print("Go And ")
		// 在这里停止，执行下一个case且只执行下一个
		fallthrough
		// 这里会报错，由于有fallthrough的存在，程序将会直接去执行下一个case结构体或者是default中的代码
		// fmt.Println("log")
	/*case "Java":
		fmt.Println(" Java.")
	case "Python":
		fmt.Println("Python.")
	*/
	default:
		fmt.Printf("%s.\n", lang2)
	}

	b()
	fmt.Println("\n")
	fmt.Println(c())

	myfunc()

	var arr []string
	arr = append(arr, "见到", "你,", "很高兴！")
	// for的遍历使用关键字 range
	// 1.只获取下标
	for key := range arr {
		fmt.Println(key)
	}
	// 2.只获取值
	for _, value := range arr {
		fmt.Println(value)
	}

}

/*
*
// if语句的表达式
*/
func pow(x, y, z float64) float64 {
	// math.Pow(x, y) 求x的y次方
	if value := math.Pow(x, y); value < z {
		return value
	}
	// return value, 报错，if语句中的变量声明只能在if语句之内使用
	return z
}

/*
*
if语句中的变量声明在连着if后面的else中使用
*/
func powElse(x, y, z float64) float64 {
	// math.Pow(x, y) 求x的y次方
	if value := math.Pow(x, y); value < z {
		return value
	} else if value == 0 {
		return value
	} else if value == 100 {
		return value
	} else {
		value = 9999
	}
	// return value, 报错，if语句中的变量声明只能在if语句之内使用
	return z
}

func Sqrt(x int) int {
	z := 1
	for x < 10 {
		// 这里的相除结果因为是int取值时会取商,而不是同float64类型时取的精确的结果值
		z -= (z*z - x) / (2 * z)
	}
	return z
}

/*
*
这里的结果会和上面不一样本质上是因为入参类型不同
*/
func SqrtFloat64(x float64) float64 {
	z := 1.0
	for x < 10 {
		// 这里的相除因为是float会取精确的结果值，而不是同int类型时取的商
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// defer 确保打开文件后能关闭该文件
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	return io.Copy(dst, src)
}

func A() {
	i := 0
	/*defer*/ fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		// 延迟打印0
		// 延迟打印1
		// 延迟打印2
		// 延迟打印3
		defer fmt.Print("延迟打印", i)
		// 按照先进后出的策略最后结果会是3210
	}
	return
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func myfunc() {
	fmt.Println("---------------------")
	i := 0
HERE:
	fmt.Println(i) // i = 0 打印出0
	i++            // 此时 i = 1
	if i < 10 {    // i = 1会进入此判断结构体
		// 跳到HERE	标签处，执行打印，程序会接着执行i++,又进入条件判断以此类推直到i=10的时候程序运行结束
		goto HERE
	}
}
