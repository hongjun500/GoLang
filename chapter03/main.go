package main

import (
	"fmt"
)

func main() {
	str := "hello go"
	// 将指针 strPointer 指向变量 str 的地址：
	strPointer := &str
	println("strPointer:", strPointer)

	var intPointer *int
	addr := &intPointer
	println("addr:", addr)

	var arrPointer []*[]string
	println("&arrPointer:", &arrPointer)

	// 定义一个名为 intParam 且类型为 int 的指针变量
	var intParam *int
	// 这里取值会报错,因为 intParam 没有初始化
	// println("对于声明的 intParam 取值:", *intParam)
	println("对于声明的 intParam 取地址:", &intParam)
	intp := 88
	fmt.Printf("指针 intParam 地址    %x\n", &intParam)
	fmt.Printf("变量 intp 地址  %x \n", &intp)
	if intParam != nil {
		fmt.Printf("%d", *intParam)
	} else {
		fmt.Println("指针 intParam 的值是空的")

		// 重新赋值指针(将指针指向变量 intp )
		intParam = &intp

		fmt.Printf("指针 intParam 指向一个实际变量后取值为%d \n", *intParam)
		fmt.Printf("指针 intParam 指向一个实际变量后指针地址为%x \n", &intParam)
		if intParam != nil {
			fmt.Println("指针 intParam 经过指向一个实际变量后不为空")
		}
	}

	// 指针指向指针示例code
	a := 1
	var ptr *int
	var pptr **int

	// 指针ptr指向a
	ptr = &a
	// *ptr = a * 100 这里改变了指针ptr的值，由于指针ptr指向了变量a,那么变量a的值也会随着ptr的改变而改变
	// 指针 pptr 指向 ptr
	pptr = &ptr
	// **pptr = *ptr * 100  同上一行，改变了 pptr 的值由于其(pptr)指向了 ptr 指针,那么 ptr 的值也改变了，指针 ptr 又指向了变量 a,因此变量 a 最后也改变了

	// a = -1  这一行改变了变量a的值，那么只要指向了这个变量a或者是间接指向了变量a的指针也会随之变量值的改变而改变
	// fmt.Printf("变量a的值=%d \n", a)
	fmt.Printf("指针变量 ptr 的值=%d \n", *ptr)
	// 取 pptr 指针的值应当使用 **varName
	fmt.Printf("指向指针变量 ptr 的指针 pptr 的值=%d \n", **pptr)

	// 将指针指向方法参数 code
	e, q := "hello", "go"
	// 再定义两个指针指向e,q变量
	// ePointer := &e
	// qPointer := &q
	fmt.Printf("没有交换值之前 ePointer 的值=%v \n", e)
	fmt.Printf("没有交换值之前 qPointer 的值=%v \n", q)
	// 之后将两个指针传入进行交换地址的值
	// 由于ePointer和qPointer分别指向了e,q变量，那么交换了ePointer,qPointer的之后就等同于将变量a,b的值进行交换
	// swap(ePointer, qPointer)

	swap(&e, &q)
	// 等同于
	/*ePointer := &e
	qPointer := &q
	swap(ePointer, qPointer)*/

	fmt.Printf("交换之后 ePointer 的值=%v \n", e)
	fmt.Printf("交换之后 qPointer 的值=%v \n", q)

	// 定义一个变量 str 并赋值 hello,go
	var strs = "hello,go"
	// 定义一个指针并指向变量 str
	var strPointers = &strs
	// 定义一个 value 变量并赋值指针 strPointer 指向变量的值
	var value = *strPointer
	fmt.Println("&strPointers", strPointers)
	fmt.Println("value", value)

	var x int = 5
	var pointerToX *int = &x
	fmt.Println("The value of x is:", x)
	fmt.Println("The address of x is:", &x)
	fmt.Println("The value of pointerToX is:", pointerToX)
	fmt.Println("The value that pointerToX points to is:", *pointerToX)

	var structPointer = new(struct {
		name string
		age  int
	})
	fmt.Println("structPointer:", *structPointer) // { 0}
}

// swap 传入两个int类型的指针参数, 将俩指针地址的值进行交换
func swap(x, y *string) {
	temp := ""
	temp = *x
	*x = *y
	*y = temp
}
