package main

import (
	"fmt"
)

func main() {

	// 指针的定义 var varName *varType
	// var intPointer *int
	// var stringPointer *string
	// 字符串数组指针
	// var strArrayPointer []*string

	// 定义一个名为intParam且类型为int的指针变量
	var intParam *int
	intp := 88
	// 打印以上两个指针的地址
	fmt.Printf("指针intParam地址    %x\n", &intParam)
	fmt.Printf("声明好的变量intp地址  %x \n", &intp)
	if intParam != nil {
		fmt.Printf("%d", *intParam)
	} else {
		fmt.Println("指针intParam的值是空的")
		fmt.Printf("指针intParam的值是%x \n", intParam)

		// 重新赋值指针(将指针指向变量intp)
		intParam = &intp

		fmt.Printf("指针intParam指向一个实际变量后取值为%d \n", *intParam)
		fmt.Printf("指针intParam指向一个实际变量后指针地址为%x \n", &intParam)
		if intParam != nil {
			fmt.Println("指针intParam经过重新赋值不再是空指针")
		}
	}

	// 将指针指向数组示例
	const MAX = 3
	arr := []int{10, 100, 1000}
	// 声明一个指针数组
	var arrays [MAX]*int
	for i := 0; i < MAX; i++ {
		// 将指针数组arrays指向数组arr
		// 此时arrays就能获取到对应arr数组的值
		arrays[i] = &arr[i]
	}

	for i := 0; i < MAX; i++ {
		// 取指针arrays所有地址的值
		fmt.Printf("arrays[%d]= %d \n", i, *arrays[i])
	}

	// 指针指向指针示例code
	a := 1
	var ptr *int
	// var varName *varType
	var pptr **int

	// 指针ptr指向a
	ptr = &a
	// *ptr = a * 100 这里改变了指针ptr的值，由于指针ptr指向了变量a,那么变量a的值也会随着ptr的改变而改变
	// 指针pptr指向ptr
	pptr = &ptr
	// **pptr = *ptr * 100  同上一行，改变了pptr的值由于其(pptr)指向了ptr指针,那么ptr的值也改变了，指针ptr又指向了变量a,因此变量a最后也改变了

	// a = -1  这一行改变了变量a的值，那么只要指向了这个变量a或者是间接指向了变量a的指针也会随之变量值的改变而改变
	// fmt.Printf("变量a的值=%d \n", a)
	fmt.Printf("指针变量ptr的值=%d \n", *ptr)
	// 取pptr指针的值应当使用**varName
	fmt.Printf("指向指针变量ptr的指针pptr的值=%d \n", **pptr)

	// 将指针指向方法参数code
	e, q := "hello", "go"
	// 再定义两个指针指向e,q变量
	// ePointer := &e
	// qPointer := &q
	fmt.Printf("没有交换值之前ePointer的值=%v \n", e)
	fmt.Printf("没有交换值之前qPointer的值=%v \n", q)
	// 之后将两个指针传入进行交换地址的值
	// 由于ePointer和qPointer分别指向了e,q变量，那么交换了ePointer,qPointer的之后就等同于将变量a,b的值进行交换
	// swap(ePointer, qPointer)

	swap(&e, &q)
	// 等同于
	/*ePointer := &e
	qPointer := &q
	swap(ePointer, qPointer)*/

	fmt.Printf("交换之后ePointer的值=%v \n", e)
	fmt.Printf("交换之后qPointer的值=%v \n", q)

}

/**
传入两个int类型的指针参数
将俩指针地址的值进行交换
*/
func swap(x, y *string) {
	temp := ""
	temp = *x
	*x = *y
	*y = temp
}
