package main

import "fmt"

func main() {
	fmt.Printf("teacher 的值为：%v \n", teacher)
	teacher.Info()

	println("teacher info 方法执行完毕")
	fmt.Printf("teacher 的值为：%v \n", teacher)

	println("------------------------------------------------------------------------")

	fmt.Printf("student 的值为：%v \n", student)
	student.Info()

	println("student info 方法执行完毕")
	fmt.Printf("student 的值为：%v \n", student)
}
