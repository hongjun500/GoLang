// 从main包开始
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	forfunc()

	fmt.Println("switch 语句部分-------------------------------------")

	// switch 语句 从上向下执行(可以没有小括号，大括号必须)
	// 同 if 一样可以有短变量声明
	// 每个 case 结构体后有默认的 break,
	// 并且只会运行表达式对应选中的 case，此时 switch 运行到此结束
	// 如果选中了 case 并且对应 case 包含一个 fallthrough 关键字语句，那么紧随其后的 case 也会执行

	// 可以没有 default 语句
	// 这个 switch 只会打印出 Go
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

	println("defer 关键字部分-------------------------------------")
	b()
	fmt.Println(" ")
	fmt.Println(c())
	println("goto 关键字部分-------------------------------------")
	myfunc()

	var arr []string
	arr = append(arr, "见到", "你,", "很高兴！")
	// for 的遍历使用关键字 range
	// 1.只获取下标
	for key := range arr {
		fmt.Println(key)
	}
	// 2.只获取值
	for _, v := range arr {
		fmt.Println(v)
	}

}

func Sqrt(x int) int {
	z := 1
	for x < 10 {
		// 这里的相除结果因为是 int 取值时会取商,而不是同 float64 类型时取的精确的结果值
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// SqrtFloat64 这里的结果会和上面不一样本质上是因为入参类型不同
func SqrtFloat64(x float64) float64 {
	z := 1.0
	for x < 10 {
		// 这里的相除因为是float会取精确的结果值，而不是同int类型时取的商
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// CopyFile 确保打开文件后能关闭该文件
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
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func myfunc() {

	i := 0
HERE:
	fmt.Println(i) // i = 0 打印出0
	i++            // 此时 i = 1
	if i < 10 {    // i = 1会进入此判断结构体
		// 跳到HERE	标签处，执行打印，程序会接着执行i++,又进入条件判断以此类推直到i=10的时候程序运行结束
		goto HERE
	}
}
