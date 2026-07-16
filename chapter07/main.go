package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("值接收者 Teacher：验证\"副本 vs 共享底层数据\"---------")
	fmt.Printf("调用前 teacher = %v\n", teacher)
	fmt.Printf("调用前 WorkAddress = %v\n", teacher.(Teacher).WorkAddress)
	teacher.Info()

	fmt.Printf("调用后 teacher = %v\n", teacher)
	fmt.Printf(" 调用后 WorkAddress = %v\n", teacher.(Teacher).WorkAddress)

	log.Println("指针接收者 *Student：验证\"副本 vs 共享底层数据\"---------")
	log.Println("------指针接收者 *Student：\"所有字段被真正改变\"---------")
	fmt.Printf("调用前 student = %v\n", student)
	student.Info()
	fmt.Printf("调用后 student = %v\n", student)

	log.Println("类型断言：判断接口底层的具体类型---------")
	describe(teacher)
	describe(student)
	assertType(teacher)
	assertType(student)


	log.Println("自定义命名类型上的方法------")
	// 自定义命名类型上的方法
	var n MyInt = 5
	fmt.Printf("MyInt(5).IsPositive() = %v\n", n.IsPositive())

	log.Println("含 nil 指针的接口 != nil -------------")
	nilInterfacePitfall()
}

// nilInterfacePitfall 演示接口的 (类型, 值) 二元组本质：
// 只有"类型和值都为 nil"时接口才 == nil。
func nilInterfacePitfall() {
	var s *Student  // 一个值为 nil 的 *Student
	var i Study = s // 装进接口后，接口记录了 (类型=*Student, 值=nil)
	// s 本身是 nil，但 i 的动态类型非 nil，所以 i != nil！
	fmt.Printf("s == nil? %v；但 i == nil? %v\n", s == nil, i == nil)
}
