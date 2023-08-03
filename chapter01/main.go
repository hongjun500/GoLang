// 从main包开始
package main

// 单个导入依赖包
import (
	"fmt"
	"log"
)
import "math"

// 分组导入多个(推荐)
import (
	"math/cmplx"
	_ "sync"
	"time"
)

// var 关键字可以出现在包或函数级别
// 声明两个string变量
var java, goLang string = "java", "go"

// 也可以使用此声明变量列表
var (
	php bool = true
	// 初始化了值可以省略类型，(自动推导类型)
	python = 1024

	price float32 = 9.9
	// 浮点类型默认float64
	price1 float64 = 9.9
)

type aliasint int

var aliasintvalue aliasint

/*
*
主程序
*/
func main() {
	// 打印不换行
	fmt.Print("hello, GoLang!")

	// 打印换行
	fmt.Println(time.Now())

	fmt.Println(time.Hour)

	fmt.Println(math.Pi)
	// 报错，小写字母开头未导出(即为无法引用)
	// 不同包内小写开头无法导入
	// fmt.Println(math.pi)

	// 换个行
	fmt.Print("\n")
	fmt.Printf("打印事先声明的变量%v,%v,%v,%v", java, goLang, python, php)
	// 换个行
	fmt.Print("\n")

	// 数值类型显示转换
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("\nx,y,f,z的值%v,%v,%v,%v \n", x, y, f, z)

	euler := 3 + 4i
	abs := cmplx.Abs(euler)
	fmt.Println("abs=", abs)

	const (
		java = iota
		golang
		python
		php
		javascript
	)
	// 可以用此方法定义“枚举类”
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tp
		pb
	)

	fmt.Println(java, golang, python, php, javascript)
	fmt.Println(b, kb, mb, gb, tp, pb)
	// iota关键字，iota初始值=0，每出现一次iota关键字直到下一个const关键字出现前iota都会都会自增1
	const (
		c0 = iota // c0=0
		c1 = iota // c1=1
		c2 = iota // c2=2
	)
	fmt.Println(c0, c1, c2)
	// 效果同上
	const (
		c3 = iota
		c4 // 省略了=iota
		c5 // 省略了=iota
	)
	fmt.Println(c3, c4, c5)

	str := "Hello, 世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Printf("%v , %c", i, ch)
	}
	fmt.Println("-----------------------")
	str = "Hello, word"
	n = len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	const (
		Big   = 1 << 2
		Small = Big >> 3
	)
	log.Print(Big)
	fmt.Println(str)

}
