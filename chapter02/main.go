// 从main包开始
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fmt.Println("计算 2 的 3 次方，如果结果小于 100 就返回结果，否则返回 100")
	value := pow(2, 3, 100)
	fmt.Println("value = ", value)

	log.Println("for 循环部分-------------------------------------")
	forExample()
	fmt.Println("switch 语句部分-------------------------------------")
	switchExample()
	log.Println("defer 关键字部分-------------------------------------")
	deferExample()

	fmt.Println("goto 关键字部分-------------------------------------")
	gotoExample()

	// CopyFile("README2.md","README.md")
}

func Sqrt(x int) int {
	z := 1
	for i:=0; i < 10; i++ {
		// 这里的相除结果因为是 int 取值时会取商,而不是同 float64 类型时取的精确的结果值
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// SqrtFloat64 这里的结果会和上面不一样本质上是因为入参类型不同
func SqrtFloat64(x float64) float64 {
	z := 1.0
	for i:=0; i < 10; i++ {
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
		err = src.Close()
		if err != nil {
			log.Fatalf("close err :%v", err.Error())
		}
	}(src)

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer func(dst *os.File) {
		err = dst.Close()
		if err != nil {
			log.Fatalf("dst.Close() err :%v", err.Error())
		}
	}(dst)

	return io.Copy(dst, src)
}

