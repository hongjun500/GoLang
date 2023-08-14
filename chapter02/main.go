// 从main包开始
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	println("计算 2 的 3 次方，如果结果小于 100 就返回结果，否则返回 100")
	value := pow(2, 3, 100)
	fmt.Println("value = ", value)

	println("for 循环部分-------------------------------------")
	forfunc()

	fmt.Println("switch 语句部分-------------------------------------")

	println("defer 关键字部分-------------------------------------")
	b()
	fmt.Println(" ")
	fmt.Println(c())
	println("goto 关键字部分-------------------------------------")
	gotoFunc()

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
